/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {
    Badge,
    BaseBehavior,
    BaseBehaviorOptions,
    ExtensionCategory,
    Graph,
    GraphData,
    GraphEvent,
    Label,
    LabelStyleProps,
    Rect,
    RectStyleProps,
    register,
    RuntimeContext
} from '@antv/g6';
import {useEffect, useRef, useState} from "react";

const DEFAULT_LEVEL = 'detailed';


const statusColors: Record<string, string> = {
    '持有权': '#17BEBB',
    '加工使用权': '#E36397',
    '经营权': '#B7AD99',
    '供给方': '#F6C86A',
};

/**
 * Draw a chart node with different ui based on the zoom level.
 */
class ChartNode extends Rect {

    get data(): Record<string, any> {
        return this.context.model.getElementDataById(this.id).data || {};
    }

    get level() {
        return this.data?.level || DEFAULT_LEVEL;
    }

    protected getLabelStyle(attributes: Required<RectStyleProps>): false | LabelStyleProps {
        return this.level === 'overview' ? {
            text: this.data?.name || '',
            fill: '#fff',
            fontSize: 20,
            fontWeight: 600,
            textAlign: "center",
            transform: 'translate(0,0)',
        } : {
            text: this.data?.name || '',
            fill: '#006fff',
            fontSize: 15,
            fontWeight: 400,
            textAlign: "left",
            transform: 'translate(-65, -15)',
        };
    }

    getKeyStyle(attributes: Required<RectStyleProps>) {
        return {
            ...super.getKeyStyle(attributes),
            fill: this.level === 'overview' ? statusColors[this.data.status] : '#fff',
        };
    }

    getPositionStyle(attributes: Required<RectStyleProps>): false | LabelStyleProps {
        if (this.level === 'overview') return false;
        return {
            text: this.data.position,
            fontSize: 9,
            fontWeight: 400,
            textTransform: 'uppercase',
            fill: '#000',
            textAlign: 'left',
            transform: 'translate(-65, 0)',
            textOverflow: 'clip',
        };
    }

    drawPositionShape(attributes: Required<RectStyleProps>, container = this) {
        const positionStyle = this.getPositionStyle(attributes);
        this.upsert('position', Label, positionStyle, container);
    }

    getStatusStyle(attributes: Required<RectStyleProps>): false | LabelStyleProps {
        if (this.level === 'overview') return false;
        return {
            text: this.data?.status,
            fontSize: 9,
            textAlign: 'left',
            transform: 'translate(40, -16)',
            padding: [0, 4],
            fill: '#fff',
            backgroundFill: statusColors[this.data?.status || ''],
        };
    }

    drawStatusShape(attributes: Required<RectStyleProps>, container = this) {
        const statusStyle = this.getStatusStyle(attributes);
        this.upsert('status', Badge, statusStyle, container);
    }

    getPhoneStyle(attributes: Required<RectStyleProps>): false | LabelStyleProps {
        if (this.level === 'overview') return false;
        return {
            text: this.data.phone,
            fontSize: 9,
            fontWeight: 300,
            textAlign: 'left',
            fill: '#000',
            transform: 'translate(-65, 20)',
        };
    }

    drawPhoneShape(attributes: Required<RectStyleProps>, container = this) {
        const style = this.getPhoneStyle(attributes);
        this.upsert('phone', Label, style, container);
    }

    render(attributes = this.parsedAttributes, container = this) {
        super.render(attributes, container);

        this.drawPositionShape(attributes, container);

        this.drawStatusShape(attributes, container);

        this.drawPhoneShape(attributes, container);
    }
}

/**
 * Implement a level of detail rendering, which will show different details based on the zoom level.
 */
class LevelOfDetail extends BaseBehavior {
    prevLevel = DEFAULT_LEVEL;
    levels = {
        ['overview']: [0, 0.6],
        ['detailed']: [0.6, Infinity],
    };

    constructor(context: RuntimeContext, options: Partial<BaseBehaviorOptions>) {
        super(context, options);
        this.bindEvents();
    }

    update(options: Partial<BaseBehaviorOptions>) {
        this.unbindEvents();
        super.update(options);
        this.bindEvents();
    }

    updateZoomLevel = async (e: any) => {
        if ('scale' in e.data) {
            const scale = e.data.scale;
            const level = Object.entries(this.levels).find(([key, [min, max]]) => scale > min && scale <= max)?.[0];
            if (level && this.prevLevel !== level) {
                const {graph} = this.context;
                graph.updateNodeData((prev) => prev.map((node) => ({...node, data: {...node.data, level}})));
                await graph.draw();
                this.prevLevel = level;
            }
        }
    };

    bindEvents() {
        const {graph} = this.context;
        graph.on(GraphEvent.AFTER_TRANSFORM, this.updateZoomLevel);
    }

    unbindEvents() {
        const {graph} = this.context;
        graph.off(GraphEvent.AFTER_TRANSFORM, this.updateZoomLevel);
    }

    destroy() {
        this.unbindEvents();
        super.destroy();
    }
}

register(ExtensionCategory.NODE, 'chart-node', ChartNode);
register(ExtensionCategory.BEHAVIOR, 'level-of-detail', LevelOfDetail);

export function TreeGraph(props: { data: GraphData }) {

    const ref = useRef<HTMLDivElement>(null);
    const [graph, setGraph] = useState<Graph>();

    useEffect(() => {
        if (!ref.current || graph) {
            return;
        }
        const instance = new Graph({
            container: ref.current,
            data: props.data,
            node: {
                type: 'chart-node',
                style: {
                    labelPlacement: 'center',
                    lineWidth: 1,
                    ports: [{placement: 'top'}, {placement: 'bottom'}],
                    radius: 2,
                    shadowBlur: 10,
                    shadowColor: '#e0e0e0',
                    shadowOffsetX: 3,
                    size: [200, 70],
                    stroke: '#C0C0C0',
                },
            },
            edge: {
                type: 'polyline',
                style: {
                    router: {
                        type: 'orth',
                    },
                    stroke: '#C0C0C0',
                },
            },
            layout: {
                type: 'dagre',
                nodeSize: [200, 70],
            },
            autoFit: 'view',
            behaviors: ['level-of-detail', 'zoom-canvas', 'drag-canvas'],
        });
        instance.render().catch(console.error);
        setGraph(instance);
    }, []);

    return (
        <div ref={ref}></div>
    )
}