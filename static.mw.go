package gin_static
import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path"
)

// INDEX 入口index.html常量
const INDEX = "index.html"

// 嵌入普通静态资源
type StaticResource struct {
	// 静态资源
	staticFS embed.FS
	// 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
	path string
}

// 静态资源被访问逻辑
func (my *StaticResource) Open(name string) (fs.File, error) {
	var fullName string
	fullName = path.Join(my.path, name)
	file, err := my.staticFS.Open(fullName)
	return file, err
}

// Static 返回静态文件中间件
func StaticEmbed(relativePath string, embedFS embed.FS) gin.HandlerFunc {
	staticRes := &StaticResource{
		staticFS: embedFS,
		path:     relativePath,
	}
	fileServer := http.FileServer(http.FS(staticRes))
	return func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}