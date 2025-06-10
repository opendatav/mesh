/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import React, {MouseEvent, useEffect, useState} from 'react';
import {useNavigate} from "react-router-dom";
import {Col, Divider, Flex, Input, Popover, Row, Tooltip} from "antd";
import {
    DownOutlined,
    EnvironmentOutlined,
    LeftCircleOutlined,
    QuestionCircleOutlined,
    RightCircleOutlined,
    RightOutlined,
    SearchOutlined,
    StarFilled,
    StarOutlined
} from "@ant-design/icons";
import {useStyles} from "@/widget";
import {Citizen} from '@ducesoft/mesh';

function PopupMenuItem(props: {
    item: Citizen,
    onClick?: (item: Citizen) => void,
    onStar?: (star: boolean, item: Citizen) => void
}) {

    const [showStar, setShowStar] = useState(props.item?.star || false);
    const {styles, cx, theme} = useStyles();

    const onHoverStar = (item: Citizen) => {
        setShowStar(!showStar);
    }

    const onClickStar = (e: MouseEvent, item: Citizen) => {
        e.stopPropagation();
        props?.onStar?.(!showStar, item);
    }

    return (
        <Flex>
            <Flex className={styles.itemView}
                  vertical={false}
                  align={'center'}
                  justify={'space-between'}
                  onMouseEnter={() => onHoverStar(props.item)}
                  onMouseLeave={() => onHoverStar(props.item)}
                  onClick={() => props?.onClick?.(props.item)}>
                <div>{props.item.name}</div>
                <div hidden={!showStar} onClick={(e: MouseEvent) => onClickStar(e, props.item)}>
                    {props.item.star ? <StarFilled/> : <StarOutlined/>}
                    &nbsp;&nbsp;
                </div>
            </Flex>
            <div>&nbsp;&nbsp;</div>
        </Flex>
    )
}

function PopupMenuTabs(props: {
    item: Citizen,
    open: boolean,
    onClick?: (plate: Citizen, item: Citizen) => void
}) {

    const {styles, cx, theme} = useStyles();
    const [open, setOpen] = useState(props.open);

    const onClickTab = (a: any) => {
        setOpen(!open);
    }

    const onClickItem = (plate: Citizen, item: Citizen) => {
        props?.onClick?.(plate, item);
    }

    return (
        <Flex vertical={true}>
            <Flex className={styles.collapseView} vertical={false} align={'center'} justify={'flex-start'}
                  onClick={onClickTab}>
                <div style={{fontSize: '18px'}}>
                    {props.item.icon}
                </div>
                <span style={{marginLeft: '10px', width: '100px'}}>{props.item.name}</span>
                {open ? <DownOutlined/> : <RightOutlined/>}
            </Flex>
            <Flex hidden={!open} vertical={true}>
                {
                    props.item?.children?.map((plate, idx) => {
                        return (
                            <div key={`${idx}`} className={styles.subsystemView}
                                 onClick={() => onClickItem(props.item, plate)}>
                                {plate.name}
                            </div>
                        )
                    })
                }
            </Flex>
        </Flex>
    )
}

export function PopupMenu(props: {
    root?: string,
    items: Citizen[],
    onClick?: (item: Citizen) => void,
    onClose?: () => void,
    className?: string
}) {

    const {styles, cx, theme} = useStyles();
    const navigate = useNavigate();
    const [activeItem, setActiveItem] = useState<Citizen>(new Citizen());
    const [selectItems, setSelectItems] = useState<Citizen[]>([]);

    const onClickTab = (plate: Citizen, item: Citizen) => {
        setActiveItem(item);
        setSelectItems([item]);
    }

    const onClickItem = (item: Citizen) => {
        navigate(Citizen.abs(item, props.root) || item.path, {replace: false});
        props?.onClose?.();
        props?.onClick?.(item);
    }

    const onStarItem = (star: boolean, item: Citizen) => {

    }

    useEffect(() => {
        setSelectItems(props.items[0]?.children || []);
    }, [props.items]);

    return (
        <Flex vertical={false} className={props?.className}>
            <Flex vertical={true} className={styles.leftView}>
                {props.items.map((item, idx) => {
                    return <PopupMenuTabs key={`${idx}`}
                                          item={item}
                                          open={idx == 0}
                                          onClick={onClickTab}/>
                })}
            </Flex>
            <Flex vertical={true} className={cx(styles.rightView, styles.scrollBar)}>
                <Flex hidden={activeItem.name != ''} vertical={true}>
                    <Flex vertical={true} align={'flex-start'} justify={'center'} style={{height: '48px'}}>
                        <div className={styles.subsystemTitle}>{activeItem.name || '子系统'}</div>
                    </Flex>
                    <Input style={{borderBottom: '1px solid #666', borderRadius: 0}} variant="borderless"
                           placeholder="请输入关键词"
                           prefix={<SearchOutlined/>}/>
                </Flex>
                {
                    selectItems.map((subsystem, idx) => {
                        return (
                            <Flex key={`${idx}`} vertical={true} align={'center'} style={{marginTop: '16px'}}>
                                <div className={styles.subsystemTitle}>{subsystem.name}</div>
                                <Row style={{width: '100%'}} gutter={[0, 0]}>
                                    {
                                        subsystem?.children?.map((plate, idx) => {
                                            return (
                                                <Col key={`${idx}`} span={8}>
                                                    <Flex vertical={true}>
                                                        <div className={styles.plateTitle}>{plate.name}</div>
                                                        {
                                                            plate?.children?.map((item, idx) => {
                                                                return (
                                                                    <PopupMenuItem key={`${idx}`}
                                                                                   item={item}
                                                                                   onClick={onClickItem}
                                                                                   onStar={onStarItem}/>
                                                                )
                                                            })
                                                        }
                                                    </Flex>
                                                </Col>
                                            )
                                        })
                                    }
                                </Row>
                            </Flex>
                        )
                    })
                }
            </Flex>
        </Flex>
    )
}

export function GridView(props: { name?: string, items: Citizen[], onClick?: (item: Citizen) => void }) {

    const {styles, cx, theme} = useStyles();

    const onClickItem = (item: Citizen) => {
        props?.onClick?.(item);
    }

    return (
        <div>
            <Flex vertical={true} align={'center'} style={{marginTop: '16px'}}>
                <Row style={{width: '100%'}} gutter={[0, 0]}>
                    {
                        props?.items?.map((item, idx) => {
                            return (
                                <Col key={`${idx}`} span={12}>
                                    <Flex vertical={true}>
                                        <div className={styles.plateTitle}>{item.name}</div>
                                        {
                                            item?.children?.map((pivot, idy) => {
                                                return (
                                                    <Flex key={`${idx}-${idy}`}
                                                          className={styles.itemView}
                                                          vertical={false}
                                                          align={'center'}
                                                          justify={'space-between'}
                                                          onClick={() => onClickItem(pivot)}>
                                                        <div>{pivot.name}</div>
                                                    </Flex>
                                                )
                                            })
                                        }
                                    </Flex>
                                </Col>
                            )
                        })
                    }
                </Row>
            </Flex>
        </div>
    )
}

export function DropdownMenu(props: { items: Citizen[], onClick?: (item: Citizen) => void }) {

    const {styles, cx, theme} = useStyles();
    const [open, setOpen] = useState(false);
    const [pivot, setPivot] = useState<Citizen>(new Citizen());

    const onHoverChange = (open: boolean) => {
        setOpen(open);
    }

    const onClickItem = (item: Citizen) => {
        setOpen(false);
        setPivot(item);
        props?.onClick?.(item);
    }

    useEffect(() => {
        setPivot(props?.items?.find((x, idx) => idx == 0)?.children?.find((x, idx) => idx == 0) || new Citizen());
    }, [props?.items])

    return (
        <div>
            <Popover placement="bottomLeft"
                     content={<GridView name={'数据中心'} items={props?.items} onClick={onClickItem}/>}
                     onOpenChange={onHoverChange} open={open}
                     arrow={false} styles={{body: {borderRadius: '0'}}}>
                <a href="#" className={styles.buttonView}>
                    <EnvironmentOutlined style={{marginRight: '8px'}}/>
                    <span style={{fontSize: '12px', lineHeight: '30px', marginRight: '5px'}}>
                        数据中心（{pivot.name}
                    </span>
                    {open ? <DownOutlined/> : <RightOutlined/>}
                </a>
            </Popover>
        </div>
    )
}

export function SideMenu(props: { root?: string, subsystem: Citizen, onClick?: (item: Citizen) => void }) {

    const {styles, cx, theme} = useStyles();
    const navigate = useNavigate();
    const [open, setOpen] = useState(true);

    const onClickItem = (item: Citizen) => {
        navigate(Citizen.abs(item, props.root) || item.path, {replace: false});
        props?.onClick?.(item);
    }

    useEffect(() => {
        for (let plate of (props.subsystem.children || [])) {
            for (let item of (plate.children || [])) {
                if (window.location.pathname && window.location.pathname.indexOf(Citizen.abs(item)) > -1) {
                    props?.onClick?.(item);
                    return;
                }
            }
        }
    }, [props.subsystem]);

    return (
        <Flex>
            <Flex hidden={!open} className={cx(styles.sideMenuView, styles.scrollBar)} vertical={true} align={'start'}>
                <div className={styles.sideMenuTitle}>{props.subsystem.name}</div>
                {
                    props.subsystem.children?.map((item, idx) => {
                        return (
                            <Flex key={`${idx}`} vertical={true} justify={'start'} style={{width: '100%'}}>
                                <div className={styles.sideMenuGroup}>
                                    <Tooltip placement="right" overlay={item.describe || ''}>
                                        <Flex vertical={false} align={'center'}>
                                            {item.name}
                                            &nbsp;
                                            <QuestionCircleOutlined/>
                                        </Flex>
                                    </Tooltip>
                                </div>
                                {
                                    item.children?.map((fn, idy) => {
                                        return (
                                            <div key={`${idx}-${idy}`}
                                                 className={styles.itemView}
                                                 onClick={() => onClickItem(fn)}>
                                                {fn.name}
                                            </div>
                                        )
                                    })
                                }
                                <Divider style={{margin: '0 0'}}/>
                            </Flex>
                        )
                    })
                }

            </Flex>
            <div style={{padding: '0.5rem 0'}} onClick={() => setOpen(!open)}>
                {open ? <LeftCircleOutlined/> : <RightCircleOutlined/>}
            </div>
        </Flex>
    );
}
