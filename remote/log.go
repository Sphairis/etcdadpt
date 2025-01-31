/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remote

import (
	"fmt"
	"time"

	"github.com/go-chassis/openlog"
)

func (c *Client) logger() openlog.Logger {
	if c.Cfg.Logger == nil {
		return openlog.GetLogger()
	}
	return c.Cfg.Logger
}

func (c *Client) logRecover(r interface{}) {
	c.logger().Error(fmt.Sprintf("embedded etcd recover: %v", r))
}

func (c *Client) logInfoOrWarn(start time.Time, message string) {
	cost := time.Since(start)
	if cost < time.Second {
		c.logger().Info(fmt.Sprintf("[%s]%s", cost, message))
		return
	}
	c.logger().Warn(fmt.Sprintf("[%s]%s", cost, message))
}

func (c *Client) logNilOrWarn(start time.Time, message string) {
	cost := time.Since(start)
	if cost < time.Second {
		return
	}
	c.logger().Warn(fmt.Sprintf("[%s]%s", cost, message))
}
