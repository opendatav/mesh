/*
 * Copyright (c) 2000, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {ProDescriptions as Describe} from '@ant-design/pro-components';
import React, {useEffect, useState} from 'react';
import {Tag} from "antd";

export function ServiceDescriber(props: { service: Record<string, any> }) {

    const [service, setService] = useState<Record<string, any>>(props.service);

    useEffect(() => {
        setService(props.service);
    }, [props.service]);

    return (
        <Describe style={{width: '30vw'}} column={1} title={`服务详情`}>
            <Describe.Item label="服务名称" valueType="text" copyable={true}>
                {service?.name}
            </Describe.Item>
            <Describe.Item label="服务提供" valueType="text">
                {service?.provider}
            </Describe.Item>
            <Describe.Item label="服务类型" valueType="text">
                {service?.type}
            </Describe.Item>
            <Describe.Item label="接口状态" valueType="text">
                <Tag color={service?.status == 'enabled' ? 'blue' : 'red'}>
                    {service?.status == 'enabled' ? '可用' : '不可用'}
                </Tag>
            </Describe.Item>
            <Describe.Item label="负载后端" valueType="jsonCode">
                {
                    ((service?.loadBalancer?.servers || []) as Record<string, any>[]).map((x, idx) => {
                        return (
                            <div style={{padding: '4px 0 4px 0', width: '20vw'}} key={idx}>
                                <Tag color={service?.serverStatus?.[x.url] == 'UP' ? 'green' : 'red'}>{x.url}</Tag>
                            </div>
                        )
                    })
                }
            </Describe.Item>
            <Describe.Item label="挂载路由" valueType="jsonCode">
                {
                    ((service?.usedBy || []) as string[]).map((x, idx) => {
                        return (
                            <div style={{padding: '4px 0 4px 0', width: '20vw'}} key={idx}>
                                <Tag key={idx} color="blue">{x}</Tag>
                            </div>
                        )
                    })
                }
            </Describe.Item>
        </Describe>
    )

}