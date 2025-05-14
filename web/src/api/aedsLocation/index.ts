import { http, jumpExport } from '@/utils/http/axios';

// 获取AED定位列表
export function List(params) {
  return http.request({
    url: '/aedsLocation/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除AED定位
export function Delete(params) {
  return http.request({
    url: '/aedsLocation/delete',
    method: 'POST',
    params,
  });
}

// 添加/编辑AED定位
export function Edit(params) {
  return http.request({
    url: '/aedsLocation/edit',
    method: 'POST',
    params,
  });
}

// 获取AED定位指定详情
export function View(params) {
  return http.request({
    url: '/aedsLocation/view',
    method: 'GET',
    params,
  });
}

// 导出AED定位
export function Export(params) {
  jumpExport('/aedsLocation/export', params);
}