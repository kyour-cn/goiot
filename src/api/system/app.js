import config from "@/config"
import http from "@/utils/request"

export default {
    list: {
        url: `${config.API_URL}/app/admin/system/appList`,
        name: "应用列表",
        get: async function (params = {}) {
            return await http.get(this.url, params);
        }
    },
    edit: {
        url: `${config.API_URL}/app/admin/system/editApp`,
        name: "新增或修改应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    },
    del: {
        url: `${config.API_URL}/app/admin/system/deleteApp`,
        name: "删除应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    }
}