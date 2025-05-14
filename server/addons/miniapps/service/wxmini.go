// Package service
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"hotgo/addons/miniapps/model"
	"hotgo/addons/miniapps/api/api"
	"hotgo/internal/consts"
	"hotgo/internal/model/input/adminin"
	"hotgo/internal/service"
	"hotgo/utility/encrypt"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// WxMiniLoginService 小程序登录服务接口
type WxMiniLoginService interface {
	// GetConfig 获取小程序配置
	GetConfig(ctx context.Context) (config *model.WxMiniAppConfig, err error)
	// Login 微信小程序登录
	Login(ctx context.Context, in *api.WxMiniLoginReq) (res *api.WxMiniLoginRes, err error)
}

// 实现接口
type sWxMiniLoginService struct{}

// 微信小程序登录服务实例
var wxMiniLoginService = &sWxMiniLoginService{}

// WxMini 获取微信小程序登录服务接口
func WxMini() WxMiniLoginService {
	return wxMiniLoginService
}

// GetConfig 获取小程序配置
func (s *sWxMiniLoginService) GetConfig(ctx context.Context) (config *model.WxMiniAppConfig, err error) {
	// 从配置文件中获取微信小程序配置
	appID := g.Cfg().MustGet(ctx, "wxmini.appid").String()
	appSecret := g.Cfg().MustGet(ctx, "wxmini.appsecret").String()

	// 使用配置文件中的微信小程序配置
	config = &model.WxMiniAppConfig{
		AppID:     appID,
		AppSecret: appSecret,
	}

	// 检测配置是否完整
	if config.AppID == "" || config.AppSecret == "" {
		return nil, gerror.New("微信小程序配置不完整，请检查配置文件中的wxmini.appid和wxmini.appsecret")
	}
	return
}

// encryptPassword 加密密码，用于登录验证
func (s *sWxMiniLoginService) encryptPassword(password string) (string, error) {
	// 先使用AES加密
	encrypted, err := encrypt.AesECBEncrypt([]byte(password), consts.RequestEncryptKey)
	if err != nil {
		return "", err
	}

	// 再使用Base64编码
	return gbase64.EncodeToString(encrypted), nil
}

// Login 微信小程序登录
func (s *sWxMiniLoginService) Login(ctx context.Context, in *api.WxMiniLoginReq) (res *api.WxMiniLoginRes, err error) {
	if err = validate.PreFilter(ctx, in); err != nil {
		return nil, err
	}

	// 获取小程序配置
	config, err := s.GetConfig(ctx)
	if err != nil {
		return nil, err
	}

	// 调用微信接口获取openid和session_key
	code2SessionURL := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.AppID, config.AppSecret, in.Code,
	)

	response, err := g.Client().Get(ctx, code2SessionURL)
	if err != nil {
		return nil, gerror.Newf("调用微信接口失败: %v", err)
	}
	defer response.Close()

	// 解析微信返回的JSON数据
	wxRes := struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}{}

	err = json.Unmarshal(response.ReadAll(), &wxRes)
	if err != nil {
		return nil, gerror.Newf("解析微信返回数据失败: %v", err)
	}

	// 检查微信返回的错误信息
	if wxRes.ErrCode != 0 {
		return nil, gerror.Newf("微信登录失败: %s (错误码: %d)", wxRes.ErrMsg, wxRes.ErrCode)
	}

	// 检查必要的返回值
	if wxRes.OpenID == "" || wxRes.SessionKey == "" {
		return nil, gerror.New("微信返回的数据不完整")
	}

	// TODO: 根据openid查询或创建用户，生成token
	// 这里简单实现，实际项目中需要根据业务需求进行调整

	// 调用主模块的登录服务生成token
	// 在实际项目中，可能需要根据openid查询用户信息，或者创建新用户
	// 这里简化处理，直接使用一个默认账号登录

	// 对密码进行加密处理
	encryptedPassword, err := s.encryptPassword("123456")
	if err != nil {
		return nil, gerror.Newf("密码加密失败: %v", err)
	}

	// 构造一个登录请求
	loginReq := &adminin.AccountLoginInp{
		Username: "admin", // 这里使用默认账号，实际项目中应该根据openid查询对应的用户
		Password: encryptedPassword,
	}

	loginRes, err := service.AdminSite().AccountLogin(ctx, loginReq)
	if err != nil {
		return nil, gerror.Newf("生成登录凭证失败: %v", err)
	}

	// 构造返回结果
	res = &api.WxMiniLoginRes{
		Token:  loginRes.Token,
		// LoginModel: loginRes,
		// OpenID:     wxRes.OpenID,
		// SessionKey: wxRes.SessionKey,
		// UnionID:    wxRes.UnionID,
	}

	return
}
