import config from "@/config"
import http from "@/utils/request"

export default {
    list: {
        url: `${config.API_URL}/admin/system/app/list`,
        name: "应用列表",
        get: async function (params = {}) {
            return await http.get(this.url, params);
        }
    },
    add: {
        url: `${config.API_URL}/admin/system/app/add`,
        name: "新增应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    },
    edit: {
        url: `${config.API_URL}/admin/system/app/edit`,
        name: "修改应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    },
    del: {
        url: `${config.API_URL}/admin/system/app/delete`,
        name: "删除应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    }
}