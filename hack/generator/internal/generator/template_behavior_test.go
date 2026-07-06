package generator

import (
	"path/filepath"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModuleTemplateFilesIncludeRouterAndServiceSystem(t *testing.T) {
	files := moduleTemplateFiles("blog")
	paths := make([]string, 0, len(files))
	for _, file := range files {
		paths = append(paths, filepath.ToSlash(file.filePath))
	}

	assert.Contains(t, paths, "modules/blog/api/blog/test.go")
	assert.Contains(t, paths, "modules/blog/controller/blog/test.go")
	assert.Contains(t, paths, "modules/blog/router/blog/router.go")
}

func TestRenderModuleTemplateRegistersHTTPModule(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "module.go.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `"devinggo/modules/blog/router/blog"`)
	assert.Contains(t, content, `modules.Register(module)`)
	assert.Contains(t, content, `s.BindHookHandler("/blog/*", ghttp.HookBeforeServe, service.Hook().BeforeServe)`)
	assert.Contains(t, content, `s.Group("/blog", func(group *ghttp.RouterGroup) {`)
	assert.Contains(t, content, `blog.BindController(group)`)
}

func TestRenderModuleRouterTemplateBindsTestController(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "router.go.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `"devinggo/modules/blog/controller/blog"`)
	assert.Contains(t, content, `blog.TestController`)
}

func TestRenderModuleTestAPITemplateUsesRelativePath(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "test_api.go.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `package blog`)
	assert.Contains(t, content, `path:"/test"`)
	assert.NotContains(t, content, `path:"/blog/test"`)
}

func TestRenderModuleTestControllerTemplateExportsController(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "test_controller.go.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `package blog`)
	assert.Contains(t, content, `TestController = testController{}`)
	assert.Contains(t, content, `blogapi "devinggo/modules/blog/api/blog"`)
	assert.Contains(t, content, `bloglogic "devinggo/modules/blog/logic/blog"`)
	assert.Contains(t, content, `message := bloglogic.New().Test(ctx)`)
}

func TestRenderModuleUpSQLTemplateEnablesModule(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "module_up_postgres.sql.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog", "date": "2026-07-06 00:00:00"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, "INSERT INTO system_modules")
	assert.Contains(t, content, "'blog'")
	assert.Contains(t, content, "installed")
	assert.Contains(t, content, "status")
}

func TestRenderServiceSystemTemplateContainsCodegenMarkers(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "module", "service_system.go.tpl"),
		g.Map{"moduleName": "blog", "moduleNameCap": "Blog"},
	)
	require.NoError(t, err)

	assert.Contains(t, content, "codegen:interfaces")
	assert.Contains(t, content, "codegen:variables")
	assert.Contains(t, content, "codegen:functions")
}

func TestRenderCRUDAPITemplateUsesSystemSharedModels(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "crud", "api.go.tpl"),
		map[string]string{
			"ModuleName":  "blog",
			"EntityName":  "BlogPost",
			"VarName":     "post",
			"PackageName": "blog",
			"ChineseName": "文章",
		},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `"devinggo/modules/system/model"`)
	assert.Contains(t, content, `"devinggo/modules/system/model/page"`)
	assert.NotContains(t, content, `"devinggo/modules/blog/model/page"`)
}

func TestRenderCRUDControllerTemplateUsesSystemBaseController(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "crud", "controller.go.tpl"),
		map[string]string{
			"ModuleName":  "blog",
			"EntityName":  "BlogPost",
			"VarName":     "post",
			"PackageName": "blog",
			"ChineseName": "文章",
		},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `"devinggo/modules/system/controller/base"`)
	assert.NotContains(t, content, `"devinggo/modules/blog/controller/base"`)
}

func TestRenderCRUDLogicTemplateUsesSystemSharedPackages(t *testing.T) {
	content, err := RenderTemplate(
		filepath.Join("..", "..", "templates", "crud", "logic.go.tpl"),
		map[string]string{
			"ModuleName":       "blog",
			"EntityName":       "BlogPost",
			"PackageName":      "blog",
			"SearchConditions": "",
			"SaveDoFields":     "",
			"UpdateDoFields":   "",
		},
	)
	require.NoError(t, err)

	assert.Contains(t, content, `"devinggo/modules/system/logic/base"`)
	assert.Contains(t, content, `"devinggo/modules/system/pkg/handler"`)
	assert.Contains(t, content, `"devinggo/modules/system/pkg/hook"`)
	assert.Contains(t, content, `"devinggo/modules/system/pkg/orm"`)
	assert.Contains(t, content, `"devinggo/modules/system/pkg/utils"`)
	assert.NotContains(t, content, `"devinggo/modules/blog/pkg/hook"`)
}

func TestRegisterRouterUsesModuleControllerPackage(t *testing.T) {
	updated, err := registerControllerInRouterContent(`package blog

func BindController(group *ghttp.RouterGroup) {
	group.Bind(
		blog.TestController,
	)
}
`, "blog", "BlogPost")
	require.NoError(t, err)

	assert.Contains(t, updated, "blog.BlogPostController")
	assert.Contains(t, updated, "blog.TestController")
}

func TestRegisterServiceInterfaceInContentUsesCodegenMarkers(t *testing.T) {
	updated, err := registerServiceInterfaceInContent(`package service

type (
	// codegen:interfaces
)

var (
	// codegen:variables
)

// codegen:functions
`, "\tIArticle interface{}\n", "\tlocalArticle IArticle\n", "func Article() IArticle { return localArticle }\n")
	require.NoError(t, err)

	assert.Contains(t, updated, "IArticle interface{}")
	assert.Contains(t, updated, "localArticle IArticle")
	assert.Contains(t, updated, "func Article() IArticle")
}
