// js api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

import { request } from '@/utils/request.js'

export default {

/**
* 获取systemModules分页列表
* @returns
*/
getPageList (params = {}) {
    return request({
        url: 'system/systemModules/index',
        method: 'get',
        params
    })
},

getList (params = {}) {
    return request({
        url: 'system/systemModules/list',
        method: 'get',
        params
    })
},

/**
* 从回收站获取systemModules数据列表
* @returns
*/
getPageRecycleList (params = {}) {
    return request({
        url: 'system/systemModules/recycle',
        method: 'get',
        params
    })
},

getRecycleList (params = {}) {
    return request({
        url: 'system/systemModules/recycleList',
        method: 'get',
        params
    })
},

/**
* 恢复systemModules数据
* @returns
*/
recoverys (data) {
    return request({
        url: 'system/systemModules/recovery',
        method: 'put',
        data
    })
},

/**
* 真实删除systemModules
* @returns
*/
realDeletes (data) {
    return request({
        url: 'system/systemModules/realDelete',
        method: 'delete',
        data
    })
},

/**
* 添加systemModules
* @returns
*/
save (data = {}) {
    return request({
        url: 'system/systemModules/save',
        method: 'post',
        data
    })
},

/**
* 更新systemModules数据
* @returns
*/
update (id, data = {}) {
    return request({
        url: 'system/systemModules/update/' + id,
        method: 'put',
        data
    })
},

/**
* 读取systemModules
* @returns
*/
read (id) {
    return request({
        url: 'system/systemModules/read/' + id,
        method: 'get'
    })
},

/**
* 将systemModules删除，有软删除则移动到回收站
* @returns
*/
deletes (data) {
    return request({
        url: 'system/systemModules/delete',
        method: 'delete',
        data
    })
},

/**
* 更改systemModules数据
* @returns
*/
changeStatus (data = {}) {
    return request({
        url: 'system/systemModules/changeStatus',
        method: 'put',
        data
    })
},

}