<template>
    <el-dialog :title="titleMap[mode]" v-model="visible" :width="500" destroy-on-close @closed="$emit('closed')">
        <el-form :model="form" :rules="rules" :disabled="mode==='show'" ref="dialogForm" label-width="100px">
            <el-form-item label="名称" prop="name">
                <el-input v-model="form.name" placeholder="请输入名称" clearable></el-input>
            </el-form-item>
            <el-form-item label="标识名称" prop="key">
                <el-input v-model="form.key" placeholder="请输入标识名称" clearable></el-input>
            </el-form-item>
            <el-form-item label="备注信息" prop="remark">
                <el-input v-model="form.remark" placeholder="请备注信息" clearable></el-input>
            </el-form-item>
            <el-form-item label="状态" prop="status">
                <el-switch v-model="form.status"></el-switch>
            </el-form-item>
<!--            <el-form-item label="图标上传" prop="status">-->
<!--                <sc-upload-->
<!--                    v-model="form.icon"-->
<!--                    :apiObj="uploadApiObj"-->
<!--                    title="请点击上传"-->
<!--                />-->
<!--            </el-form-item>-->
        </el-form>
        <template #footer>
            <el-button @click="visible=false">取 消</el-button>
            <el-button v-if="mode!=='show'" type="primary" :loading="isSaveing" @click="submit()">保 存</el-button>
        </template>
    </el-dialog>
</template>

<script>
import productApi from "@/api/platform/product"
export default {
    emits: ['success', 'closed'],
    data() {
        return {
            mode: "add",
            titleMap: {
                add: '新增',
                edit: '编辑',
                show: '查看'
            },
            visible: false,
            isSaveing: false,
            //表单数据
            form: {
                name: "",
                remark: "",
                key: "",
                status: true,
                sort: 0,
                icon: ""
            },
            //验证规则
            rules: {
                name: [
                    {required: true, message: '请输入姓名'},
                ]
            },
            //所需数据选项
            groups: [],
            groupsProps: {
                value: "id",
                emitPath: false,
                checkStrictly: true
            },
            // uploadApiObj: this.$API.common.upload.uploadImage
        }
    },
    methods: {
        //显示
        open(mode = 'add') {
            this.mode = mode;
            this.visible = true;
            return this
        },
        //表单提交方法
        async submit() {
            const valid = await this.$refs.dialogForm.validate()
            if (!valid) {
                return false
            }
            this.isSaveing = true;
            const data = this.form
            data.status = data.status ? 1 : 0;
            const res = await productApi[this.mode].post(data);
            this.isSaveing = false;
            this.visible = false;
            if (res.code === 0) {
                this.$emit('success', this.form, this.mode)
                this.$message.success("操作成功")
            } else {
                await this.$alert(res.message, "提示", {type: 'error'})
            }
        },
        //表单注入数据
        setData(data) {
            // this.form.id = data.id
            // this.form.name = data.name
            // this.form.remark = data.remark
            // this.form.status = data.status === 1
            // this.form.sort = data.sort
            // this.form.icon = data.icon
            //可以和上面一样单个注入，也可以像下面一样直接合并进去
            Object.assign(this.form, data)
            this.form.status = data.status === 1
        }
    }
}
</script>

<style>
</style>
