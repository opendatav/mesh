/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {Avatar, Timeline} from "antd";
import {ClockCircleOutlined} from "@ant-design/icons";
import React from "react";
import {CheckCard} from "@ant-design/pro-components";
import moment from "moment/moment";

export interface HistoryClock {
    level: string;
    timestamp?: Date;
    title: string
    memo: string;
}

function Clock(props: { clock: HistoryClock }) {
    return (
        <div>
            <CheckCard title={props.clock.title}
                       avatar={<Avatar src={null} size="large"/>}
                       description={<span>{props.clock.memo}</span>}
                       value={''}
            />
        </div>
    )
}

export function HistoryViewer(props: { clocks: HistoryClock[] }) {

    return (
        <div>
            <Timeline mode={'left'} style={{padding: '20px 0 0 20px', width: '240px'}}
                      items={props.clocks.map((clock, idx) => {
                          return {
                              key: `${idx}`,
                              dot: <ClockCircleOutlined style={{fontSize: '18px'}}/>,
                              label: moment(clock.timestamp).format('YYYY-MM-DD HH:mm:ss'),
                              children: <Clock clock={clock}/>,
                              color: clock.level === '4' ? 'red' : 'green'
                          }
                      })}
            />
        </div>
    )
}
