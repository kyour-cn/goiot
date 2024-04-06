<template>
    <el-dialog title="设置参数" v-model="visible" :width="900" destroy-on-close @closed="emits('closed')">
        <el-row>
            <el-col :span="10">
                <el-button @click="addField">添加</el-button>
                <sc-table :data="fieldList" :hide-pagination="true" :hide-do="true" :height="400">
                    <el-table-column label="参数名称" prop="name"></el-table-column>
                    <el-table-column label="类型" prop="type">
                        <template #default="scope">
                            {{getType(scope.row.type)}}
                        </template>
                    </el-table-column>
                    <el-table-column label="操作">
                        <template #default="scope" >
                            <el-button link @click="edit(scope.row)">编辑</el-button>
                            <el-button link @click="del(scope)">删除</el-button>
                        </template>
                    </el-table-column>
                </sc-table>
            </el-col>
            <el-col :span="14">
                <el-form :model="fieldItem" :rules="rules" ref="dialogForm" label-width="100px" style="height: 100%">
                    <el-form-item label="名称" prop="name">
                        <el-input v-model="fieldItem.name" placeholder="请输入名称" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="类型" prop="extend_field">
                        <el-radio-group v-model="fieldItem.type" class="ml-4" @change="changeType">
                            <el-radio v-for="item in fieldType" :label="item.value" size="large">{{item.name}}</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item v-if="fieldItem.type === 'radio' || fieldItem.type === 'checkbox'" label="选项" prop="option">
                        <div style="width: 12px"></div>
                        <el-button v-for="(item, key) in fieldItem.option" size="small" @click="delOption(key)" style="margin-top: 10px">{{item}}</el-button>
                        <el-button type="success" size="small" @click="addOption" style="margin-top: 10px">+添加</el-button>
                    </el-form-item>
                    <el-form-item label="默认值" prop="default">
                        <el-input v-model="fieldItem.default" placeholder="请输入默认值" clearable></el-input>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>

        <template #footer>
            <el-button @click="visible=false">取 消</el-button>
            <el-button type="primary" @click="submit()">保 存</el-button>
        </template>
    </el-dialog>

    <el-dialog v-model="addOptionVisible" title="添加选项" :width="250">
        <el-form>
            <el-input v-model="optionValue" autocomplete="off" />
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="confirmOption">
          确认
        </el-button>
      </span>
        </template>
    </el-dialog>
</template>

<script setup>
import ScTable from "@/components/scTable/index.vue";
import {getCurrentInstance, ref} from "vue";

const proxy = getCurrentInstance().proxy

const emits = defineEmits([
    'closed',
    'success'
])

const visible = ref(false);

const fieldType = ref([
    {
        name: '文本',
        value: 'text'
    },
    {
        name: '单选',
        value: 'radio'
    },
    {
        name: '多选',
        value: 'checkbox'
    },
    {
        name: '日期',
        value: 'date'
    },
    {
        name: '日期时间',
        value: 'datetime'
    }
]);

const fieldList = ref([]);

const addOptionVisible = ref(false);
const optionValue = ref('');

const addField = () => {
    fieldList.value.push({
        name: '未命名',
        type: 'text',
        remark: ''
    })
    fieldItem.value = fieldList.value[fieldList.value.length - 1];
}

const edit = (row) => {
    fieldItem.value = row;
    visible.value = true;
}

const del = (row) => {
    fieldList.value.splice(row.$index, 1)
}

const fieldItem = ref({
    name: '',
    type: 'text',
    remark: '',
    option: []
});

// 改变类型事件
const changeType = (value) => {
    if (fieldItem.value.option === undefined && (value === 'radio' || value === 'checkbox')) {
        fieldItem.value.option = [];
    }
}

// 添加选项
const addOption = () => {
    addOptionVisible.value = true;
}

// 确认添加选项
const confirmOption = () => {
    fieldItem.value.option.push(optionValue.value);
    optionValue.value = '';
    addOptionVisible.value = false;
}

const delOption = (index) => {
    fieldItem.value.option.splice(index, 1);
}

const getType = (type) => {
    return fieldType.value.find(item => item.value === type).name;
}

const rules = {
    name: [
        {required: true, message: '请输入名称'},
    ]
}

const submit = async () => {
    // const form = ref(dialogForm.value);
    // form.value.validate((valid) => {
    //     if (valid) {
    //         emits('success', form.value.model);
    //         visible.value = false;
    //     } else {
    //         return false;
    //     }
    // });

    const res = await proxy.$API.platform.category.edit.post({
        id: id,
        extend_field: JSON.stringify(fieldList.value)
    });
    if (res.code === 0) {
        emits('success', fieldItem.value);
        proxy.$message.success("操作成功")
        visible.value = false;
    } else {
        await proxy.$alert(res.message, "提示", {type: 'error'})
    }

}

let id = 0;

defineExpose({
    open(){
        visible.value = true
        return this
    },
    setData(data){
        id = data.id;
        if (data.extend_field) {
            fieldList.value = JSON.parse(data.extend_field);
        }
    }
})

</script>

<style>
</style>
