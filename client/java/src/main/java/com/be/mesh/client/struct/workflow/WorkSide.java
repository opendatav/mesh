/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */
package com.be.mesh.client.struct.workflow;

import com.be.mesh.client.annotate.Index;
import lombok.Data;

import java.io.Serializable;

/**
 * @author coyzeng@gmail.com
 */
@Data
public class WorkSide implements Serializable {

    private static final long serialVersionUID = 32241564753916622L;
    /**
     * Workflow side src name
     */
    @Index(0)
    private String src;
    /**
     * Workflow side dst name
     */
    @Index(1)
    private String dst;
    /**
     * Workflow side condition
     * <pre>
     *     func Invoke(ctx context.Context, ctx *types.WorkContext) (string, error) {
     * 	       return "true", nil
     *     }
     * </pre>
     */
    @Index(2)
    private String condition;
}
