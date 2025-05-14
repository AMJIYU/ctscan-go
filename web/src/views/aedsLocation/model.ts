import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';

export class State {
  public id = 0; // id
  public longitude = null; // 经度
  public latitude = null; // 纬度
  public address = ''; // 地址
  public createdAt = ''; // created_at
  public updatedAt = ''; // updated_at

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  longitude: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入经度',
  },
  latitude: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入纬度',
  },
  address: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入地址',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'address',
    component: 'NInput',
    label: '地址',
    componentProps: {
      placeholder: '请输入地址',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: 'created_at',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'longitude',
    component: 'NInputNumber',
    label: '经度',
    componentProps: {
      placeholder: '请输入经度',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'latitude',
    component: 'NInputNumber',
    label: '纬度',
    componentProps: {
      placeholder: '请输入纬度',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  
  {
    field: 'id',
    component: 'NInputNumber',
    label: 'id',
    componentProps: {
      placeholder: '请输入id',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
 
]);

// 表格列
export const columns = [
  {
    title: 'id',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '经度',
    key: 'longitude',
    align: 'left',
    width: -1,
  },
  {
    title: '纬度',
    key: 'latitude',
    align: 'left',
    width: -1,
  },
  {
    title: '地址',
    key: 'address',
    align: 'left',
    width: -1,
  },
  {
    title: 'created_at',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: 'updated_at',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
];