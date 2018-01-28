// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "./common"
import "strconv"

// nolint
func init() {
	common.Template_topic_alt_handle = Template_topic_alt
	common.Ctemplates = append(common.Ctemplates,"topic_alt")
	common.TmplPtrMap["topic_alt"] = &common.Template_topic_alt_handle
	common.TmplPtrMap["o_topic_alt"] = Template_topic_alt
}

// nolint
func Template_topic_alt(tmpl_topic_alt_vars common.TopicPage, w http.ResponseWriter) error {
w.Write(header_0)
w.Write([]byte(tmpl_topic_alt_vars.Title))
w.Write(header_1)
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.Name))
w.Write(header_2)
w.Write([]byte(tmpl_topic_alt_vars.Header.Theme.Name))
w.Write(header_3)
if len(tmpl_topic_alt_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Stylesheets {
w.Write(header_4)
w.Write([]byte(item))
w.Write(header_5)
}
}
w.Write(header_6)
if len(tmpl_topic_alt_vars.Header.Scripts) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Scripts {
w.Write(header_7)
w.Write([]byte(item))
w.Write(header_8)
}
}
w.Write(header_9)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(header_10)
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.URL))
w.Write(header_11)
if tmpl_topic_alt_vars.Header.MetaDesc != "" {
w.Write(header_12)
w.Write([]byte(tmpl_topic_alt_vars.Header.MetaDesc))
w.Write(header_13)
}
w.Write(header_14)
if !tmpl_topic_alt_vars.CurrentUser.IsSuperMod {
w.Write(header_15)
}
w.Write(header_16)
w.Write(menu_0)
w.Write(menu_1)
w.Write([]byte(tmpl_topic_alt_vars.Header.Site.ShortName))
w.Write(menu_2)
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
w.Write(menu_3)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Link))
w.Write(menu_4)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(menu_5)
} else {
w.Write(menu_6)
}
w.Write(menu_7)
w.Write(header_17)
if tmpl_topic_alt_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_18)
}
w.Write(header_19)
if len(tmpl_topic_alt_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.NoticeList {
w.Write(header_20)
w.Write([]byte(item))
w.Write(header_21)
}
}
if tmpl_topic_alt_vars.Page > 1 {
w.Write(topic_alt_0)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_1)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page - 1)))
w.Write(topic_alt_2)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_3)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page - 1)))
w.Write(topic_alt_4)
}
if tmpl_topic_alt_vars.LastPage != tmpl_topic_alt_vars.Page {
w.Write(topic_alt_5)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_6)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page + 1)))
w.Write(topic_alt_7)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_8)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Page + 1)))
w.Write(topic_alt_9)
}
w.Write(topic_alt_10)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_11)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_12)
if tmpl_topic_alt_vars.Topic.Sticky {
w.Write(topic_alt_13)
} else {
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_14)
}
}
w.Write(topic_alt_15)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Title))
w.Write(topic_alt_16)
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_17)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_alt_18)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Title))
w.Write(topic_alt_19)
}
w.Write(topic_alt_20)
if tmpl_topic_alt_vars.Poll.ID > 0 {
w.Write(topic_alt_21)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_22)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_23)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_24)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Avatar))
w.Write(topic_alt_25)
w.Write([]byte(tmpl_topic_alt_vars.Topic.UserLink))
w.Write(topic_alt_26)
w.Write([]byte(tmpl_topic_alt_vars.Topic.CreatedByName))
w.Write(topic_alt_27)
if tmpl_topic_alt_vars.Topic.Tag != "" {
w.Write(topic_alt_28)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Tag))
w.Write(topic_alt_29)
} else {
w.Write(topic_alt_30)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.Level)))
w.Write(topic_alt_31)
}
w.Write(topic_alt_32)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_33)
if len(tmpl_topic_alt_vars.Poll.QuickOptions) != 0 {
for _, item := range tmpl_topic_alt_vars.Poll.QuickOptions {
w.Write(topic_alt_34)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_35)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_36)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_37)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_38)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_39)
w.Write([]byte(item.Value))
w.Write(topic_alt_40)
}
}
w.Write(topic_alt_41)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_42)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_43)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Poll.ID)))
w.Write(topic_alt_44)
}
w.Write(topic_alt_45)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Avatar))
w.Write(topic_alt_46)
w.Write([]byte(tmpl_topic_alt_vars.Topic.UserLink))
w.Write(topic_alt_47)
w.Write([]byte(tmpl_topic_alt_vars.Topic.CreatedByName))
w.Write(topic_alt_48)
if tmpl_topic_alt_vars.Topic.Tag != "" {
w.Write(topic_alt_49)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Tag))
w.Write(topic_alt_50)
} else {
w.Write(topic_alt_51)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.Level)))
w.Write(topic_alt_52)
}
w.Write(topic_alt_53)
w.Write([]byte(tmpl_topic_alt_vars.Topic.ContentHTML))
w.Write(topic_alt_54)
w.Write([]byte(tmpl_topic_alt_vars.Topic.Content))
w.Write(topic_alt_55)
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
if tmpl_topic_alt_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_alt_56)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_57)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_58)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_alt_59)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_60)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.DeleteTopic {
w.Write(topic_alt_61)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_62)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_63)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.CloseTopic {
if tmpl_topic_alt_vars.Topic.IsClosed {
w.Write(topic_alt_64)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_65)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_66)
} else {
w.Write(topic_alt_67)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_68)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_69)
}
}
if tmpl_topic_alt_vars.CurrentUser.Perms.PinTopic {
if tmpl_topic_alt_vars.Topic.Sticky {
w.Write(topic_alt_70)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_71)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_72)
} else {
w.Write(topic_alt_73)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_74)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_75)
}
}
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_76)
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_77)
}
w.Write(topic_alt_78)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_79)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_80)
}
w.Write(topic_alt_81)
if tmpl_topic_alt_vars.Topic.LikeCount > 0 {
w.Write(topic_alt_82)
}
w.Write(topic_alt_83)
if tmpl_topic_alt_vars.Topic.LikeCount > 0 {
w.Write(topic_alt_84)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.LikeCount)))
w.Write(topic_alt_85)
}
w.Write(topic_alt_86)
w.Write([]byte(tmpl_topic_alt_vars.Topic.RelativeCreatedAt))
w.Write(topic_alt_87)
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_88)
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_89)
w.Write([]byte(tmpl_topic_alt_vars.Topic.IPAddress))
w.Write(topic_alt_90)
}
w.Write(topic_alt_91)
if len(tmpl_topic_alt_vars.ItemList) != 0 {
for _, item := range tmpl_topic_alt_vars.ItemList {
w.Write(topic_alt_92)
if item.ActionType != "" {
w.Write(topic_alt_93)
}
w.Write(topic_alt_94)
w.Write([]byte(item.Avatar))
w.Write(topic_alt_95)
w.Write([]byte(item.UserLink))
w.Write(topic_alt_96)
w.Write([]byte(item.CreatedByName))
w.Write(topic_alt_97)
if item.Tag != "" {
w.Write(topic_alt_98)
w.Write([]byte(item.Tag))
w.Write(topic_alt_99)
} else {
w.Write(topic_alt_100)
w.Write([]byte(strconv.Itoa(item.Level)))
w.Write(topic_alt_101)
}
w.Write(topic_alt_102)
if item.ActionType != "" {
w.Write(topic_alt_103)
}
w.Write(topic_alt_104)
if item.ActionType != "" {
w.Write(topic_alt_105)
w.Write([]byte(item.ActionIcon))
w.Write(topic_alt_106)
w.Write([]byte(item.ActionType))
w.Write(topic_alt_107)
} else {
w.Write(topic_alt_108)
w.Write([]byte(item.ContentHtml))
w.Write(topic_alt_109)
if tmpl_topic_alt_vars.CurrentUser.Loggedin {
if tmpl_topic_alt_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_alt_110)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_111)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_112)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.EditReply {
w.Write(topic_alt_113)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_114)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_115)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.DeleteReply {
w.Write(topic_alt_116)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_117)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_118)
}
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_119)
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_120)
}
w.Write(topic_alt_121)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_alt_122)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_123)
}
w.Write(topic_alt_124)
if item.LikeCount > 0 {
w.Write(topic_alt_125)
}
w.Write(topic_alt_126)
if item.LikeCount > 0 {
w.Write(topic_alt_127)
w.Write([]byte(strconv.Itoa(item.LikeCount)))
w.Write(topic_alt_128)
}
w.Write(topic_alt_129)
w.Write([]byte(item.RelativeCreatedAt))
w.Write(topic_alt_130)
if tmpl_topic_alt_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_alt_131)
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_132)
w.Write([]byte(item.IPAddress))
w.Write(topic_alt_133)
}
w.Write(topic_alt_134)
}
w.Write(topic_alt_135)
}
}
w.Write(topic_alt_136)
if tmpl_topic_alt_vars.CurrentUser.Perms.CreateReply {
w.Write(topic_alt_137)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Avatar))
w.Write(topic_alt_138)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Link))
w.Write(topic_alt_139)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Name))
w.Write(topic_alt_140)
if tmpl_topic_alt_vars.CurrentUser.Tag != "" {
w.Write(topic_alt_141)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Tag))
w.Write(topic_alt_142)
} else {
w.Write(topic_alt_143)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.CurrentUser.Level)))
w.Write(topic_alt_144)
}
w.Write(topic_alt_145)
w.Write([]byte(tmpl_topic_alt_vars.CurrentUser.Session))
w.Write(topic_alt_146)
w.Write([]byte(strconv.Itoa(tmpl_topic_alt_vars.Topic.ID)))
w.Write(topic_alt_147)
if tmpl_topic_alt_vars.CurrentUser.Perms.UploadFiles {
w.Write(topic_alt_148)
}
w.Write(topic_alt_149)
}
w.Write(topic_alt_150)
w.Write(footer_0)
w.Write([]byte(common.BuildWidget("footer",tmpl_topic_alt_vars.Header)))
w.Write(footer_1)
if len(tmpl_topic_alt_vars.Header.Themes) != 0 {
for _, item := range tmpl_topic_alt_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_2)
w.Write([]byte(item.Name))
w.Write(footer_3)
if tmpl_topic_alt_vars.Header.Theme.Name == item.Name {
w.Write(footer_4)
}
w.Write(footer_5)
w.Write([]byte(item.FriendlyName))
w.Write(footer_6)
}
}
}
w.Write(footer_7)
w.Write([]byte(common.BuildWidget("rightSidebar",tmpl_topic_alt_vars.Header)))
w.Write(footer_8)
	return nil
}
