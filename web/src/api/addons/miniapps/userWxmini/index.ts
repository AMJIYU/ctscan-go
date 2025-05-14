import { http, jumpExport } from '@/utils/http/axios';

// 获取微信小程序用户表列表
export function List(params) {
  return http.request({
    url: '/miniapps/userWxmini/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除微信小程序用户表
export function Delete(params) {
  return http.request({
    url: '/miniapps/userWxmini/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑微信小程序用户表
export function Edit(params) {
  return http.request({
    url: '/miniapps/userWxmini/edit',
    method: 'POST',
    params,
  });
}

// 获取微信小程序用户表指定详情
export function View(params) {
  return http.request({
    url: '/miniapps/userWxmini/view',
    method: 'GET',
    params,
  });
}

// 导出微信小程序用户表
export function Export(params) {
  jumpExport('/miniapps/userWxmini/export', params);
}