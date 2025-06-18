package extract_task

import (
	"context"
	"errors"

	"github.com/SnakeHacker/deepkg/admin/common"
	"github.com/SnakeHacker/deepkg/admin/internal/dao"
	"github.com/SnakeHacker/deepkg/admin/internal/model/gorm_model"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"
	"github.com/SnakeHacker/deepkg/admin/internal/types"
	"github.com/golang/glog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExtractTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExtractTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExtractTaskLogic {
	return &GetExtractTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExtractTaskLogic) GetExtractTask(req *types.GetExtractTaskReq) (resp *types.GetExtractTaskResp, err error) {

	taskModel, err := dao.SelectExtractTaskByID(l.svcCtx.DB, int(req.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	creator, err := dao.SelectUserByID(l.svcCtx.DB, int64(taskModel.CreatorID))
	if err != nil {
		glog.Error(err)
		return
	}

	extractTaskDocModels, err := dao.SelectExtractTaskDocuments(l.svcCtx.DB, int(taskModel.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	docIDs := []int64{}
	for _, extractTaskDocModel := range extractTaskDocModels {
		docIDs = append(docIDs, int64(extractTaskDocModel.DocID))
	}

	docModels, err := dao.SelectDocumentModelByIDs(l.svcCtx.DB, docIDs)
	if err != nil {
		glog.Error(err)
		return
	}

	docs := []types.Document{}
	for _, docModel := range docModels {
		docs = append(docs, types.Document{
			ID:      int64(docModel.ID),
			DocName: docModel.DocName,
			DocDesc: docModel.DocDesc,
		})
	}

	extractTaskTripleModels, err := dao.SelectExtractTaskTriples(l.svcCtx.DB, int(taskModel.ID))
	if err != nil {
		glog.Error(err)
		return
	}

	tripleIDs := []int64{}
	for _, extractTaskTripleModel := range extractTaskTripleModels {
		tripleIDs = append(tripleIDs, int64(extractTaskTripleModel.TripleID))
	}

	tripleModels, err := dao.SelectSchemaTriplesByIDs(l.svcCtx.DB, tripleIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	ontologyIDs := []int64{}
	for _, tripleModel := range tripleModels {
		ontologyIDs = append(ontologyIDs, int64(tripleModel.SourceOntologyID), int64(tripleModel.TargetOntologyID))
	}

	ontologyModels, err := dao.SelectSchemaOntologiesByIDs(l.svcCtx.DB, ontologyIDs)
	if err != nil {
		glog.Error(err)
		return
	}
	ontologuMap := map[int64]gorm_model.SchemaOntology{}
	for _, ontologyModel := range ontologyModels {
		ontologuMap[int64(ontologyModel.ID)] = *ontologyModel
	}

	triples := []types.SchemaTriple{}
	for _, tripleModel := range tripleModels {
		sourceOntology, ok := ontologuMap[int64(tripleModel.SourceOntologyID)]
		if !ok {
			err = errors.New("source ontology not found")
			glog.Error(err)
			return
		}

		targetOntology, ok := ontologuMap[int64(tripleModel.TargetOntologyID)]
		if !ok {
			err = errors.New("target ontology not found")
			glog.Error(err)
			return
		}

		triples = append(triples, types.SchemaTriple{
			ID:                 int64(tripleModel.ID),
			SourceOntologyID:   int64(tripleModel.SourceOntologyID),
			TargetOntologyID:   int64(tripleModel.TargetOntologyID),
			SourceOntologyName: sourceOntology.OntologyName,
			TargetOntologyName: targetOntology.OntologyName,
			Relationship:       tripleModel.Relationship,
			WorkSpaceID:        int64(tripleModel.WorkSpaceID),
		})
	}

	task := types.ExtractTask{
		ID:          int64(taskModel.ID),
		TaskName:    taskModel.TaskName,
		Remark:      taskModel.Remark,
		TaskStatus:  taskModel.TaskStatus,
		WorkSpaceID: int64(taskModel.WorkSpaceID),
		Published:   taskModel.Published,
		Docs:        docs,
		Triples:     triples,

		CreatorID:   int64(creator.ID),
		CreatorName: creator.Username,
		CreatedAt:   taskModel.CreatedAt.Format(common.TIME_FORMAT),
	}

	resp = &types.GetExtractTaskResp{
		ExtractTask: task,
	}

	return
}
