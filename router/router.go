package router

import (
	"work_order/apis/tpl"
	"work_order/pkg/jwtauth"
	"work_order/router/dashboard"
	"work_order/router/process"
	systemRouter "work_order/router/system"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// 静态文件
	sysStaticFileRouter(g, r)

	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)

	// 无需认证
	systemRouter.SysNoCheckRoleRouter(g)

	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysStaticFileRouter(r *gin.RouterGroup, g *gin.Engine) {
	r.Static("/static", "./static")
	g.LoadHTMLGlob("template/web/index.html")
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	v1 := r.Group("/api/v1")

	// 兼容前后端不分离的情
	r.GET("/", tpl.Tpl)

	// 首页
	dashboard.RegisterDashboardRouter(v1, authMiddleware)

	// 系统管理
	systemRouter.RegisterPageRouter(v1, authMiddleware)       // 页面路由
	systemRouter.RegisterBaseRouter(v1, authMiddleware)       // 基本路由
	systemRouter.RegisterDeptRouter(v1, authMiddleware)       // 部门路由
	systemRouter.RegisterSysUserRouter(v1, authMiddleware)    // 系统用户路由
	systemRouter.RegisterRoleRouter(v1, authMiddleware)       // 角色路由
	systemRouter.RegisterUserCenterRouter(v1, authMiddleware) // 用户中心路由
	systemRouter.RegisterPostRouter(v1, authMiddleware)       // 注册后路由
	systemRouter.RegisterMenuRouter(v1, authMiddleware)       // 菜单路由
	systemRouter.RegisterLoginLogRouter(v1, authMiddleware)   // 登录路由
	systemRouter.RegisterSysSettingRouter(v1, authMiddleware) //系统设置路由

	// 流程中心
	process.RegisterClassifyRouter(v1, authMiddleware)  // 分类路由
	process.RegisterProcessRouter(v1, authMiddleware)   //进程路由
	process.RegisterTaskRouter(v1, authMiddleware)      // 任务路由
	process.RegisterTplRouter(v1, authMiddleware)       //模版路由
	process.RegisterWorkOrderRouter(v1, authMiddleware) // 工单路由
}
