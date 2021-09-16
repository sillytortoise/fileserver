package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type FileInfo struct {
	Name    string `json:"name"`
	Modtime string `json:"modtime"`
	Size    int    `json:"size"`
	Folder  bool   `json:"folder"`
}

type Node struct {
	Value    string `json:"value"`
	Label    string `json:"label"`
	Children []Node `json:"children"`
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) GetFiles() {
	dirname := "files"
	var files []FileInfo
	fs, _ := ioutil.ReadDir(dirname)
	for _, fi := range fs {
		var f FileInfo
		f.Name = fi.Name()
		f.Modtime = fi.ModTime().Local().String()
		f.Size = int(fi.Size() / 1024) //KB
		if !fi.IsDir() {
			f.Folder = false
		} else {
			f.Folder = true
		}
		files = append(files, f)

	}
	c.Data["json"] = files
	c.ServeJSON()
}

func (c *MainController) Delete() {
	path := c.Ctx.Input.Param(":splat")
	fmt.Println(path)
	err := os.Remove("files/" + path)
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString("success")
}

func (c *MainController) Upload() {
	f, h, _ := c.GetFile("file") //获取上传的文件
	//创建目录
	var uploadDir string
	if c.Ctx.Input.Param(":splat") == "" {
		uploadDir = "files/"
	} else {
		uploadDir = "files/" + c.Ctx.Input.Param(":splat") + "/"
	}
	fileName := h.Filename
	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("file", fpath)
	c.Ctx.WriteString("success")
}

func (c *MainController) Download() {
	path := "files/" + c.Ctx.Input.Param(":splat")
	fmt.Println(path)
	c.Ctx.Output.Download(path)
}

func (c *MainController) GetFolderFiles() {
	path := c.Ctx.Input.Param(":splat")
	folderpath := "files" + "/" + path
	var files []FileInfo
	fs, err := ioutil.ReadDir(folderpath)
	if err != nil {
		c.Ctx.WriteString("error")
		return
	}
	for _, fi := range fs {
		var f FileInfo
		f.Name = fi.Name()
		f.Modtime = fi.ModTime().Local().String()
		f.Size = int(fi.Size() / 1024) //KB
		if !fi.IsDir() {
			f.Folder = false
		} else {
			f.Folder = true
		}
		files = append(files, f)
	}
	c.Data["json"] = files
	c.ServeJSON()
}

func (c *MainController) IsEmpty() {
	path := c.Ctx.Input.Param(":splat")
	folderpath := "files/" + path
	fs, err := ioutil.ReadDir(folderpath)
	if err != nil {
		panic(err)
	}
	if len(fs) == 0 {
		c.Ctx.WriteString("true")
	} else {
		c.Ctx.WriteString("false")
	}
}

func (c *MainController) CreateFolder() {
	folder := c.Ctx.Input.Param(":splat")
	err := os.Mkdir("files/"+folder, 0766)
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString("success")
}

func (c *MainController) GetDir() {
	nodelist := getChildren("files")
	c.Data["json"] = nodelist
	c.ServeJSON()
}

func getChildren(path string) []Node {
	fs, _ := ioutil.ReadDir(path)
	if len(fs) == 0 {
		return nil
	}
	var nodelist []Node
	for _, fi := range fs {
		var node Node

		if fi.IsDir() {
			node.Label = fi.Name()
			node.Value = fi.Name()
			node.Children = getChildren(path + "/" + fi.Name())
			nodelist = append(nodelist, node)
		}
	}
	if len(nodelist) > 0 {
		return nodelist
	}
	return nil
}

func (c *MainController) GetStaticFile() {
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./views" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *MainController) Movefile() {
	dest := c.GetString("dest")
	file := c.GetString("file")
	parts := strings.Split(file, "/")
	filenme := parts[len(parts)-1]
	if dest == "/" {
		dest = ""
	}
	err := os.Rename("files"+file, "files"+dest+"/"+filenme)
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString("")
}
