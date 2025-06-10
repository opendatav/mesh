/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {ModalForm, ProForm, ProFormDateRangePicker, ProFormSelect, ProFormText,} from '@ant-design/pro-components';
import {message} from 'antd';

export function RouteEditor(props: { visible: boolean, onCall: (open: boolean) => void }) {

    return (
        <ModalForm
            title="新增网络路由"
            open={props?.visible}
            onFinish={async () => {
                message.success('提交成功');
                return true;
            }}
            onOpenChange={props?.onCall}
        >
            <ProForm.Group>
                <ProFormText
                    width="md"
                    name="name"
                    label="签约客户名称"
                    tooltip="最长为 24 位"
                    placeholder="请输入名称"
                />

                <ProFormText
                    width="md"
                    name="company"
                    label="我方公司名称"
                    placeholder="请输入名称"
                />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText
                    width="md"
                    name="contract"
                    label="合同名称"
                    placeholder="请输入名称"
                />
                <ProFormDateRangePicker name="contractTime" label="合同生效时间"/>
            </ProForm.Group>
            <ProForm.Group>
                <ProFormSelect
                    options={[
                        {
                            value: 'chapter',
                            label: '盖章后生效',
                        },
                    ]}
                    width="xs"
                    name="useMode"
                    label="合同约定生效方式"
                />
                <ProFormSelect
                    width="xs"
                    options={[
                        {
                            value: 'time',
                            label: '履行完终止',
                        },
                    ]}
                    name="unusedMode"
                    label="合同约定失效效方式"
                />
            </ProForm.Group>
            <ProFormText width="sm" name="id" label="主合同编号"/>
            <ProFormText
                name="project"
                disabled
                label="项目名称"
                initialValue="xxxx项目"
            />
            <ProFormText
                width="xs"
                name="mangerName"
                disabled
                label="商务经理"
                initialValue="启途"
            />
        </ModalForm>
    );
};