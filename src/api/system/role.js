import config from "@/config"
import http from "@/utils/request"

export default {
    list: {
        url: `${config.API_URL}/app/admin/system/roleList`,
        name: "获取角色列表",
        get: async function (params) {
            return await http.get(this.url, params);
        }
    },
    edit: {
        url: `${config.API_URL}/app/admin/system/editRole`,
        name: "新增或修改角色",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    },
    del: {
        url: `${config.API_URL}/app/admin/system/deleteRole`,
        name: "删除应用",
        post: async function (data = {}) {
            return await http.post(this.url, data);
        }
    }
}