import config from "@/config";
import http from "@/utils/request";

export default {
    list: {
        url: `${config.API_URL}/admin/platform/device/list`,
        name: '设备列表',
        get: async function (data, config = {}) {
            return await http.get(this.url, data, config);
        }
    },
    add: {
        url: `${config.API_URL}/admin/platform/device/add`,
        name: '添加设备',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
    edit: {
        url: `${config.API_URL}/admin/platform/device/edit`,
        name: '编辑设备',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
    delete: {
        url: `${config.API_URL}/admin/platform/device/delete`,
        name: '删除设备',
        post: async function (data, config = {}) {
            return await http.post(this.url, data, config);
        }
    },
}