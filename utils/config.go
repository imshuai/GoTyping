package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DBAddress           string `json:"db_address"`
	DBPort              int    `json:"db_port"`
	DBUser              string `json:"db_user"`
	DBPassword          string `json:"db_passwd"`
	DBName              string `json:"db_name"`
	DBCharset           string `json:"charset"`
	DBIdleConnectionNum int    `json:"db_idle_connection_num"`
	DBMaxConnectionNum  int    `json:"db_max_connectiong_num"`
}

func (c *DatabaseConfig) String() string {
	return c.DBUser + ":" +
		c.DBPassword + "@tcp(" +
		c.DBAddress + ":" +
		strconv.Itoa(c.DBPort) + ")/" +
		c.DBName + "?charset=" +
		c.DBCharset + "&timeout=5s&parseTime=True&loc=Asia%2FChongqing"
}

//Check 检查数据库配置文件内容是否正确
func (c *DatabaseConfig) Check() error {
	if c.DBAddress == "" {
		return errors.New("invalid database address")
	}
	if c.DBCharset == "" {
		return errors.New("invalid database charset")
	}
	if c.DBName == "" {
		return errors.New("invalid database name")
	}
	if c.DBPassword == "" {
		return errors.New("invalid database password")
	}
	if c.DBPort > 65536 || c.DBPort < 1 {
		return errors.New("invalid database port")
	}
	if c.DBUser == "" {
		return errors.New("invalid database username")
	}
	return nil
}

func parseDatabaseConfig() (cfg *DatabaseConfig, err error) {
	byts, err := ioutil.ReadFile("db_config.1.json") //开发配置
	if err != nil {
		log.Fatalln("read database config file db_config.json fail with error:", err)
		return nil, err
	}
	err = json.Unmarshal(byts, cfg)
	if err != nil {
		log.Fatalln("read database config file db_config.json fail with error:", err)
		return nil, err
	}
	if err = cfg.Check(); err != nil {
		return nil, err
	}
	return
}

type keywords []string

func (k keywords) MarshalText() (text []byte, err error) {
	if len(k) == 0 {
		return []byte(""), nil
	}
	s := strings.Join(k, ",")
	return []byte(s), nil
}

func (k keywords) UnmarshalText(text []byte) error {
	k = k[:0]
	s := strings.Split(string(text), ",")
	k = append(k, s...)
	return nil
}

//SiteConfig 网站配置
type SiteConfig struct {
	SiteName             string   `json:"site_name"`
	SiteDescription      string   `json:"site_description"`
	SiteKeyWords         keywords `json:"site_keywords"`
	SiteOwner            string   `json:"site_owner"`
	SiteOwnerDescription string   `json:"site_owner_description"`
	SiteOwnerWeibo       string   `json:"site_owner_weibo"`
	SiteOwnerWechat      string   `json:"site_owner_wechat"`
	SiteOwnerGithub      string   `json:"site_owner_github"`
}
