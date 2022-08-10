package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Post_20220810_090735 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20220810_090735{}
	m.Created = "20220810_090735"

	migration.Register("Post_20220810_090735", m)
}

// Run the migrations
func (m *Post_20220810_090735) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE post(`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(128) NOT NULL,`body` longtext  NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Post_20220810_090735) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `post`")
}
