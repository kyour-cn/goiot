import config from "@/config"
import http from "@/utils/request"

export default {
    list: {
        url: `${config.API_URL}/app/admin/system/ruleList`,
        name: "获取权限",
        get: async function (params = {}) {
            return await http.get(this.url, params);
        }
    },
    edit: {
        url: `${config.API_URL}/app/admin/system/editRule`,
        name: "编辑权限",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    },
    delete: {
        url: `${config.API_URL}/app/admin/system/deleteRule`,
        name: "删除权限",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    }
}