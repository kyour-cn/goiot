import config from "@/config";
import http from "@/utils/request";

export default {
    list: {
        url: `${config.API_URL}/admin/platform/product/list`,
        name: '产品列表',
        get: async function (data, config = {}) {
            return await http.get(this.url, data, config);
        }
    },
    add: {
        url: `${config.API_URL}/admin/platform/product/add`,
        name: '添加产品',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
    edit: {
        url: `${config.API_URL}/admin/platform/product/edit`,
        name: '编辑产品',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
    delete: {
        url: `${config.API_URL}/admin/platform/product/delete`,
        name: '删除产品',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
}