package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var gorerr error

func init() {
	gormDB, gorerr = gorm.Open(mysql.Open("p2p:p2pA!123@tcp(192.168.66.149:3306)/zjbxinsurance?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if gorerr != nil {
		panic("failed to connect database")
	}
}

// 1. User 模型（用户）：一对多关联 Post（一个用户发布多篇文章）
type User struct {
	ID        uint           `gorm:"primaryKey;comment:用户ID"`                           // 主键（默认自增）
	Username  string         `gorm:"size:50;not null;unique;comment:用户名"`               // 唯一非空，避免重复
	Sex       string         `gorm:"size:100;not null;comment:性别"`                      // 性别
	PostCount uint           `gorm:"size:100;comment:文章数量"`                             // 文章数量
	CreatedAt time.Time      `gorm:"comment:创建时间"`                                      // Gorm 自动维护：记录创建时间
	UpdatedAt time.Time      `gorm:"comment:更新时间"`                                      // Gorm 自动维护：记录更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`                               // 软删除（非物理删除），需加 index 优化查询
	Posts     []Post         `gorm:"foreignKey:UserID;references:ID;comment:用户发布的所有文章"` // 一对多关联：通过 Post 的 UserID 关联 User 的 ID
}

// 2. Post 模型（文章）：多对一关联 User + 一对多关联 Comment
type Post struct {
	ID            uint           `gorm:"primaryKey;comment:文章ID"`
	Content       string         `gorm:"type:text;not null;comment:文章内容"`                // 长文本类型，存储文章正文
	CommentStatus uint           `gorm:"size:100;comment:评论状态 1 为有评论 0为无评论"`             // 评论状态 1 为有评论 0为无评论
	UserID        uint           `gorm:"not null;comment:关联的用户ID（外键）"`                   // 外键：关联 User 的 ID，非空（文章必须有作者）
	User          User           `gorm:"foreignKey:UserID;references:ID;comment:文章所属用户"` // 多对一关联：通过 UserID 关联 User
	CreatedAt     time.Time      `gorm:"comment:创建时间"`
	UpdatedAt     time.Time      `gorm:"comment:更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"index;comment:软删除标记"`
	Comments      []Comment      `gorm:"foreignKey:PostID;references:ID;comment:文章的所有评论"` // 一对多关联：通过 Comment 的 PostID 关联 Post 的 ID
}

// 3. Comment 模型（评论）：多对一关联 Post
type Comment struct {
	ID        uint           `gorm:"primaryKey;comment:评论ID"`
	Content   string         `gorm:"size:500;not null;comment:评论内容"`                 // 评论长度限制，避免垃圾内容
	PostID    uint           `gorm:"not null;comment:关联的文章ID（外键）"`                   // 外键：关联 Post 的 ID，非空（评论必须属于某篇文章）
	Post      Post           `gorm:"foreignKey:PostID;references:ID;comment:评论所属文章"` // 多对一关联：通过 PostID 关联 Post
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`
}

func main() {
	// gorm模型定义
	gormDB.AutoMigrate(&User{}, &Post{}, &Comment{})

	// // 数据准备
	gormDB.Create(&User{Username: "张三", Sex: "男", PostCount: 2})
	gormDB.Create(&User{Username: "李四", Sex: "女", PostCount: 2})

	gormDB.Create(&Post{Content: "GORM 入门1", UserID: 1, CommentStatus: 1})
	gormDB.Create(&Post{Content: "GORM 入门2", UserID: 1, CommentStatus: 1})

	gormDB.Create(&Post{Content: "GORM 进阶1", UserID: 2, CommentStatus: 1})
	gormDB.Create(&Post{Content: "GORM 进阶2", UserID: 2, CommentStatus: 1})

	gormDB.Create(&Comment{Content: "ORM1 入门的评论", PostID: 1})
	gormDB.Create(&Comment{Content: "ORM1 入门的评论", PostID: 1})

	gormDB.Create(&Comment{Content: "ORM2 入门的评论", PostID: 1})
	gormDB.Create(&Comment{Content: "ORM2 入门的评论", PostID: 2})

	gormDB.Create(&Comment{Content: "GORM1 进阶的评论", PostID: 3})
	gormDB.Create(&Comment{Content: "GORM1 进阶的评论", PostID: 3})

	gormDB.Create(&Comment{Content: "GORM1 进阶的评论", PostID: 3})
	gormDB.Create(&Comment{Content: "GORM2 进阶的评论", PostID: 4})

	// // 关联查询
	// //使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// var posts []Post
	// gormDB.Debug().Preload("Comments").Where("user_id = ?", 1).Find(&posts)
	// jsonByte, _ := json.Marshal(posts)
	// fmt.Println(string(jsonByte))

	// // 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	// var post Post
	// gormDB.Debug().Model(&Comment{}).Select("*").Group("post_id").Order("count(post_id) desc").Limit(1).Find(&post)
	// jsonByte, _ = json.Marshal(post)
	// fmt.Println(string(jsonByte))

	// // 新增文章，自动更新用户文章数量
	// gormDB.Create(&Post{Content: "GORM 入门3", UserID: 1, CommentStatus: 1})

	// //删除评论，自动更新文章评论状态
	// var comment Comment
	// gormDB.Debug().Where("id = ?", 8).Find(&comment)
	// gormDB.Debug().Where("id = ?", comment.ID).Delete(&comment)

}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 更新文章评论状态  如果没有评论了，则设置为0 无评论状态
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", 0)
		fmt.Println("删除评论后，更新文章评论状态为无评论")
	} else {
		fmt.Println("删除评论后，文章仍有评论", count)
	}
	return
}
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户文章数量
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1"))
	return
}
