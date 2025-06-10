/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {Divider, List} from "antd";
import {useStyles} from "@/widget";
import {InfoCircleOutlined} from "@ant-design/icons";

export function MailBox() {

    const {styles, cx, theme} = useStyles();
    const data = [
        {
            title: '蓝象智联登记技术服务商，资质待审核',
            timestamp: '2024年08月01日 12:00:00',
            status: 1
        },
        {
            title: '中国移动登记技术服务商，资质待审核',
            timestamp: '2024年08月01日 12:00:00',
            status: 2
        },
        {
            title: '中国电信登记技术服务商，资质待审核',
            timestamp: '2024年08月01日 12:00:00',
            status: 4
        },
        {
            title: '工商银行登记技术服务商，资质待审核',
            timestamp: '2024年08月01日 12:00:00',
            status: 4
        },
    ];

    return (
        <div style={{width: '300px', margin: '10px'}}>
            <div className={styles.plateTitle}>
                未读消息(300条)
            </div>
            <Divider style={{margin: '0'}}/>
            <List
                itemLayout="horizontal"
                dataSource={data}
                renderItem={(item, idx) => (
                    <List.Item>
                        <List.Item.Meta
                            title={
                                <a style={{fontSize: '12px'}} className={styles.logonIdText} href="#">
                                    &nbsp;{item.title}&nbsp;&nbsp;
                                    <InfoCircleOutlined style={{color: 'green'}}/>
                                </a>}
                            description={<span className={styles.nameText}>{item.timestamp}</span>}
                        />
                    </List.Item>
                )}
            />
        </div>
    )
}