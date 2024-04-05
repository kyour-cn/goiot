<template>
	<el-dialog
		v-model="state.visible"
		destroy-on-close
		:title="state.titleMap[state.mode]"
		:width="500"
		@closed="$emit('closed')"
	>
		<el-form
			ref="dialogForm"
			:disabled="state.mode === 'show'"
			:model="state.form"
			:rules="state.rules"
			label-position="left"
			label-width="100px"
		>
			<el-form-item label="角色名称" prop="label">
				<el-input v-model="state.form.name" clearable></el-input>
			</el-form-item>
			<el-form-item label="排序" prop="sort">
				<el-input-number
					v-model="state.form.sort"
					controls-position="right"
					style="width: 100%;"
					:min="1"
				/>
			</el-form-item>
			<el-form-item label="是否有效" prop="status">
				<el-switch v-model="state.form.status" :active-value="true" :inactive-value="false"></el-switch>
			</el-form-item>
			<el-form-item label="备注" prop="remark">
				<el-input v-model="state.form.remark" clearable type="textarea"></el-input>
			</el-form-item>
		</el-form>
		<template #footer>
			<el-button @click="state.visible=false">取 消</el-button>
			<el-button v-if="state.mode !== 'show'" :loading="state.isSaving" type="primary" @click="submit()">保 存</el-button>
		</template>
	</el-dialog>
</template>

<script setup>
import roleApi from "@/api/system/role";

import {getCurrentInstance, reactive, ref} from "vue";

const {proxy} = getCurrentInstance()
const emits = defineEmits(['success', 'closed'])
const dialogForm = ref(null)

const state = reactive({
    mode: "add",
    titleMap: {
        add: '新增',
        edit: '编辑',
        show: '查看'
    },
    visible: false,
    isSaving: false,
    //表单数据
    form: {
        id: "",
        name: "",
        sort: 1,
        status: 1,
        remark: "",
        app_id: 0
    },
    //验证规则
    rules: {
        sort: [
            {required: true, message: '请输入排序', trigger: 'change'}
        ],
        name: [
            {required: true, message: '请输入角色名称'}
        ]
    }
})

//表单提交方法
const submit = () => {
    dialogForm.value.validate(async (valid) => {
        if (valid) {
            state.isSaving = true;
            
            const data = state.form
            data.status = data.status ? 1 : 0;
            
            const res = await roleApi[state.mode].post(data);
            state.isSaving = false;
            if (res.code === 0) {
                emits('success', state.form, state.mode)
                state.visible = false;
                proxy.$message.success("操作成功")
            } else {
                await proxy.$alert(res.message, "提示", {type: 'error'})
            }
        }
    })
}

const open = (mode = 'add') => {
    state.mode = mode
    state.visible = true
}

const setData = (data) => {
    Object.assign(state.form, data)
    state.form.status = data.status === 1
}

//暴露给父组件的方法
defineExpose({
    open, setData
})
//
// export default {
// 	emits: ['success', 'closed'],
// 	props: ['selectedApp'],
//
// 	methods: {
// 		//显示
// 		open(mode = 'add') {
// 			this.mode = mode;
// 			this.visible = true;
// 			return this
// 		},
// 		//表单提交方法
// 		submit() {
// 			this.$refs.dialogForm.validate(async (valid) => {
// 				if (!valid) {
// 					return false
// 				}
// 				this.isSaving = true;
// 				const data = this.form
// 				data.status = data.status ? 1 : 0;
// 				const res = await this.$API.system.role.edit.post(data);
// 				this.isSaving = false;
// 				this.visible = false;
// 				if (res.code === 0) {
// 					this.$emit('success', this.form, this.mode)
// 					this.$message.success("操作成功")
// 				} else {
// 					this.$alert(res.message, "提示", {type: 'error'})
// 				}
// 			})
// 		},
// 		//表单注入数据
// 		setData(data) {
// 			//可以和上面一样单个注入，也可以像下面一样直接合并进去
// 			Object.assign(this.form, data)
//
// 			this.form.status = data.status === 1
// 		}
// 	}
// }
</script>
