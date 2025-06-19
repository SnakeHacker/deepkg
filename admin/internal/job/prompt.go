package job

import (
	"fmt"

	"github.com/SnakeHacker/deepkg/admin/common/knowledge_graph"
)

func ExtractEntityPrompt(content string, ontologies []string) string {
	result := fmt.Sprintf(`
语料如下:
%s


已知有本体概念如下:
%v

请从语料中根据本体概念抽取实体，已存在的实体请不要重复抽取。以JSON格式返回
	{
		"entities": [
			{
				"entity": "实体1",
				"type": "本体概念1"
			},
			{
				"entity": "实体2",
				"type": "本体概念2"
			}
		]
	}
}
`, content, ontologies)
	return result
}

func ExtractPropsPrompt(content string, entity string, props []string) string {
	result := fmt.Sprintf(`
语料如下:
%s


已知有实体如下:
%v
实体拥有的属性范围如下:
%v

请从语料中根据实体抽取对应的实体属性和属性值， 属性范围请限定在上述提供的属性中， 如果不存在相关属性则忽略。以JSON格式返回
{
	"props": [
		{
			"prop": "属性1",
			"value": "属性值1"
		},
		{
			"prop": "属性2",
			"value": "属性值2"
		}
	]
}
`, content, entity, props)
	return result
}

func ExtractRelationshipPrompt(content string, entities []knowledge_graph.Entity, tripleContent []string) string {

	entitiesInfo := ""
	for _, entity := range entities {
		entitiesInfo += fmt.Sprintf("{实体: %s, 本体: %s} \n", entity.EntityName, entity.Type)
	}

	tripleInfo := ""
	for _, triple := range tripleContent {
		tripleInfo += fmt.Sprintf("%s \n", triple)
	}
	result := fmt.Sprintf(`
语料如下:
%s

已知有实体如下:
%v

实体是本体的实例化，已知有三元组概念（源本体1、关系、目标本体2）如下:
%v

请从语料中根据已有的实体抽取对应的三元组，三元组范围请限定在上述提供的三元组中，如果不存在则忽略。以JSON格式返回
{
	"relationships": [
			{
				"source": "源实体1",
				"rel": "关系1",
				"target": "目标实体1"
			},
			{
				"source": "源实体2",
				"rel": "关系2",
				"target": "目标实体2"
			}
		]
	}
}
`, content, entitiesInfo, tripleInfo)
	return result
}
