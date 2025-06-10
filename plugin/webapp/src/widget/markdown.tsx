/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import ReactMarkdown from "react-markdown";
import RemarkMath from "remark-math";
import RemarkBreaks from "remark-breaks";
import RehypeKatex from "rehype-katex";
import RemarkGfm from "remark-gfm";
import RehypeHighlight from "rehype-highlight";
import React, {RefObject, useMemo} from "react";
import "katex/dist/katex.min.css";

function escapeDollarNumber(text: string) {
    let escapedText = "";

    for (let i = 0; i < text.length; i += 1) {
        let char = text[i];
        const nextChar = text[i + 1] || " ";

        if (char === "$" && nextChar >= "0" && nextChar <= "9") {
            char = "\\$";
        }

        escapedText += char;
    }

    return escapedText;
}

function escapeBrackets(text: string) {
    const pattern =
        /(```[\s\S]*?```|`.*?`)|\\\[([\s\S]*?[^\\])\\\]|\\\((.*?)\\\)/g;
    return text.replace(
        pattern,
        (match, codeBlock, squareBracket, roundBracket) => {
            if (codeBlock) {
                return codeBlock;
            } else if (squareBracket) {
                return `$$${squareBracket}$$`;
            } else if (roundBracket) {
                return `$${roundBracket}$`;
            }
            return match;
        },
    );
}

function tryWrapHtmlCode(text: string) {
    // try add wrap html code (fixed: html codeblock include 2 newline)
    return text
        .replace(
            /([`]*?)(\w*?)([\n\r]*?)(<!DOCTYPE html>)/g,
            (match, quoteStart, lang, newLine, doctype) => {
                return !quoteStart ? "\n```html\n" + doctype : match;
            },
        )
        .replace(
            /(<\/body>)([\r\n\s]*?)(<\/html>)([\n\r]*?)([`]*?)([\n\r]*?)/g,
            (match, bodyEnd, space, htmlEnd, newLine, quoteEnd) => {
                return !quoteEnd ? bodyEnd + space + htmlEnd + "\n```\n" : match;
            },
        );
}

export function Markdown(props: {
                             children: string;
                             loading?: boolean;
                             fontSize?: number;
                             fontFamily?: string;
                             parentRef?: RefObject<HTMLDivElement>;
                             defaultShow?: boolean;
                             split?: string;
                             style?: {};
                         } & React.DOMAttributes<HTMLDivElement>,
) {
    const escapedContent = useMemo(() => {
        if (!props.split) {
            return tryWrapHtmlCode(escapeBrackets(escapeDollarNumber(props.children)));
        }
        for (let paragraph of props.children.split('---')) {
            const texts = paragraph.trim().split('\n');
            if (texts[0].replaceAll('#', '').trim() == props.split) {
                return tryWrapHtmlCode(escapeBrackets(escapeDollarNumber(texts.filter((x, idx) => idx > 0).join('\n'))));
            }
        }
        return '待编写';
    }, [props.children, props.split]);

    return (
        <div
            style={Object.assign({
                fontSize: `${props.fontSize ?? 14}px`,
                fontFamily: props.fontFamily || "inherit",
            }, props.style || {})}
            onContextMenu={props.onContextMenu}
            onDoubleClickCapture={props.onDoubleClickCapture}
            dir="auto"
        >
            <ReactMarkdown
                remarkPlugins={[RemarkMath, RemarkGfm, RemarkBreaks]}
                rehypePlugins={[
                    RehypeKatex,
                    [
                        RehypeHighlight,
                        {
                            detect: false,
                            ignoreMissing: true,
                        },
                    ],
                ]}
            >
                {escapedContent}
            </ReactMarkdown>
        </div>
    );
}