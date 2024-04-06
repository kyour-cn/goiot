<template>
    <el-container>
        <el-header>
            <div class="left-panel">
                <el-button type="primary" icon="el-icon-plus" @click="add"></el-button>
                <el-button type="danger" plain icon="el-icon-delete" :disabled="selection.length === 0"
                           @click="batch_del"></el-button>
            </div>
        </el-header>
        <el-main class="nopadding">
            <sc-table ref="tableRef" :apiObj="list.apiObj" row-key="id" @selection-change="selectionChange" stripe>
                <el-table-column type="selection" width="50"></el-table-column>
<!--                <el-table-column label="图标" prop="status" width="60">-->
<!--                    <template #default="scope">-->
<!--                        <img :src="scope.row.icon" alt="" style="width: 25px;height: 25px">-->
<!--                    </template>-->
<!--                </el-table-column>-->
                <el-table-column label="名称" prop="name" width="180"></el-table-column>
                <el-table-column label="产品标识" prop="key" width="180"></el-table-column>
                <el-table-column label="备注信息" prop="remark"></el-table-column>
                <el-table-column label="状态" prop="status" width="60">
                    <template #default="scope">
                        <sc-status-indicator v-if="scope.row.status" type="success"></sc-status-indicator>
                        <sc-status-indicator v-if="!scope.row.status" pulse type="danger"></sc-status-indicator>
                    </template>
                </el-table-column>
                <el-table-column label="操作" fixed="right" align="right" width="300">
                    <template #default="scope">
                        <el-button-group>
                            <el-button text type="primary" plain size="small" @click="tableEdit(scope.row)">编辑</el-button>
                            <el-button text type="primary" plain size="small" @click="fieldEdit(scope.row)">参数设置</el-button>
                            <el-popconfirm title="确定删除吗？" @confirm="tableDel(scope.row, scope.$index)">
                                <template #reference>
                                    <el-button text plain type="danger" size="small">删除</el-button>
                                </template>
                            </el-popconfirm>
                        </el-button-group>
                    </template>
                </el-table-column>
            </sc-table>
        </el-main>
    </el-container>

    <save-dialog
        v-if="state.save"
        ref="saveDialogRef"
        @success="handleSaveSuccess"
        @closed="state.save=false"
    />

    <field-dialog
        v-if="state.field"
        ref="fieldDialogRef"
        @success="handleSaveSuccess"
        @closed="state.field=false"
    />

</template>

<script setup name="product">
import SaveDialog from './save.vue'
import FieldDialog from './field.vue'
import ScStatusIndicator from "@/components/scMini/scStatusIndicator.vue";
import ScTable from "@/components/scTable/index.vue";
import productApi from "@/api/platform/product"
import {getCurrentInstance, nextTick, reactive, ref} from 'vue'

const proxy = getCurrentInstance().proxy

const state = reactive({
    save: false,
    field: false
})

const list = reactive({
    apiObj: productApi.list
})

const selection = ref([])

const saveDialogRef = ref(null)
const fieldDialogRef = ref(null)
const tableRef = ref(null)

const add = () => {
    state.save = true
    nextTick(() => {
        saveDialogRef.value.open()
    })
}

const tableEdit = (row) => {
    state.save = true
    nextTick(() => {
        saveDialogRef.value.open('edit').setData(row)
    })
}

const fieldEdit = (row) => {
    state.field = true
    nextTick(() => {
        fieldDialogRef.value.open().setData(row)
    })
}

const tableDel = async (row, index) => {
    const reqData = {id: row.id};
    const res = await productApi.delete.post(reqData);
    if (res.code === 0) {
        tableRef.value.refresh()
        proxy.$message.success("删除成功")
    } else {
        await proxy.$alert(res.message, "提示", {type: 'error'})
    }
}

const batch_del = async () => {
    const confirmRes = await proxy.$confirm(`确定删除选中的 ${selection.value.length} 项吗？`, '提示', {
        type: 'warning',
        confirmButtonText: '删除',
        confirmButtonClass: 'el-button--danger'
    }).catch(() => {});

    if (!confirmRes) {
        return false
    }

    const ids = selection.value.map(v => v.id);
    const res = await productApi.delete.post({id: ids});
    if (res.code === 0) {
        tableRef.value.refresh()
        proxy.$message.success("操作成功")
    } else {
        await proxy.$alert(res.message, "提示", {type: 'error'})
    }
}

const selectionChange = (val) => {
    selection.value = val
}

const handleSaveSuccess = (data, mode) => {
    tableRef.value.refresh()
}

</script>
