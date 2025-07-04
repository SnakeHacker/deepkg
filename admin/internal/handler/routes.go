// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "github.com/SnakeHacker/deepkg/admin/internal/handler/admin"
	chat "github.com/SnakeHacker/deepkg/admin/internal/handler/chat"
	document "github.com/SnakeHacker/deepkg/admin/internal/handler/document"
	document_dir "github.com/SnakeHacker/deepkg/admin/internal/handler/document_dir"
	extract_task "github.com/SnakeHacker/deepkg/admin/internal/handler/extract_task"
	extract_task_result "github.com/SnakeHacker/deepkg/admin/internal/handler/extract_task_result"
	file "github.com/SnakeHacker/deepkg/admin/internal/handler/file"
	knowledge_graph_workspace "github.com/SnakeHacker/deepkg/admin/internal/handler/knowledge_graph_workspace"
	org "github.com/SnakeHacker/deepkg/admin/internal/handler/org"
	schema_ontology "github.com/SnakeHacker/deepkg/admin/internal/handler/schema_ontology"
	schema_ontology_prop "github.com/SnakeHacker/deepkg/admin/internal/handler/schema_ontology_prop"
	schema_triple "github.com/SnakeHacker/deepkg/admin/internal/handler/schema_triple"
	session "github.com/SnakeHacker/deepkg/admin/internal/handler/session"
	user "github.com/SnakeHacker/deepkg/admin/internal/handler/user"
	"github.com/SnakeHacker/deepkg/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: admin.HealthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/chat",
				Handler: chat.StreamChatHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/document/get",
					Handler: document.GetDocumentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document/list",
					Handler: document.GetDocumentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document/create",
					Handler: document.CreateDocumentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document/update",
					Handler: document.UpdateDocumentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document/delete",
					Handler: document.DeleteDocumentsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/document_dir/get",
					Handler: document_dir.GetDocumentDirHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document_dir/list",
					Handler: document_dir.GetDocumentDirListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document_dir/create",
					Handler: document_dir.CreateDocumentDirHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document_dir/update",
					Handler: document_dir.UpdateDocumentDirHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/document_dir/delete",
					Handler: document_dir.DeleteDocumentDirsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: file.UploadFileHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/org/get",
					Handler: org.GetOrgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/org/list",
					Handler: org.GetOrgListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/org/create",
					Handler: org.CreateOrgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/org/update",
					Handler: org.UpdateOrgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/org/delete",
					Handler: org.DeleteOrgsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology/get",
					Handler: schema_ontology.GetSchemaOntologyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology/list",
					Handler: schema_ontology.GetSchemaOntologyListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology/create",
					Handler: schema_ontology.CreateSchemaOntologyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology/update",
					Handler: schema_ontology.UpdateSchemaOntologyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology/delete",
					Handler: schema_ontology.DeleteSchemaOntologysHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology_prop/get",
					Handler: schema_ontology_prop.GetSchemaOntologyPropHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology_prop/list",
					Handler: schema_ontology_prop.GetSchemaOntologyPropListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology_prop/create",
					Handler: schema_ontology_prop.CreateSchemaOntologyPropHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology_prop/update",
					Handler: schema_ontology_prop.UpdateSchemaOntologyPropHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_ontology_prop/delete",
					Handler: schema_ontology_prop.DeleteSchemaOntologyPropsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/schema_triple/get",
					Handler: schema_triple.GetSchemaTripleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_triple/list",
					Handler: schema_triple.GetSchemaTripleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_triple/create",
					Handler: schema_triple.CreateSchemaTripleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_triple/update",
					Handler: schema_triple.UpdateSchemaTripleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema_triple/delete",
					Handler: schema_triple.DeleteSchemaTriplesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/session/login",
				Handler: session.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/session/captcha",
				Handler: session.GetCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/session/publickey",
				Handler: session.GetPublicKeyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/session/logout",
					Handler: session.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/get",
					Handler: user.GetUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUsersHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/get",
					Handler: extract_task.GetExtractTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/list",
					Handler: extract_task.GetExtractTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/create",
					Handler: extract_task.CreateExtractTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/update",
					Handler: extract_task.UpdateExtractTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/delete",
					Handler: extract_task.DeleteExtractTasksHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/publish",
					Handler: extract_task.PublishExtractTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task/run",
					Handler: extract_task.RunExtractTaskHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/entity/list",
					Handler: extract_task_result.GetEntityListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/entity/create",
					Handler: extract_task_result.CreateEntityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/entity/update",
					Handler: extract_task_result.UpdateEntityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/entity/delete",
					Handler: extract_task_result.DeleteEntitiesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/prop/list",
					Handler: extract_task_result.GetPropListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/prop/create",
					Handler: extract_task_result.CreatePropHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/prop/update",
					Handler: extract_task_result.UpdatePropHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/prop/delete",
					Handler: extract_task_result.DeletePropsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/relationship/list",
					Handler: extract_task_result.GetRelationshipListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/relationship/create",
					Handler: extract_task_result.CreateRelationshipHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/relationship/update",
					Handler: extract_task_result.UpdateRelationshipHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/relationship/delete",
					Handler: extract_task_result.DeleteRelationshipsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/extract_task_result/get",
					Handler: extract_task_result.GetExtractTaskResultHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtX},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/knowledge_graph_workspace/get",
					Handler: knowledge_graph_workspace.GetKnowledgeGraphWorkspaceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/knowledge_graph_workspace/list",
					Handler: knowledge_graph_workspace.GetKnowledgeGraphWorkspaceListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/knowledge_graph_workspace/create",
					Handler: knowledge_graph_workspace.CreateKnowledgeGraphWorkspaceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/knowledge_graph_workspace/update",
					Handler: knowledge_graph_workspace.UpdateKnowledgeGraphWorkspaceHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/knowledge_graph_workspace/delete",
					Handler: knowledge_graph_workspace.DeleteKnowledgeGraphWorkspacesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)
}
