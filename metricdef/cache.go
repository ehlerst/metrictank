/*
 * Copyright (c) 2015, Raintank Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metricdef

import (
	"github.com/ctdk/goas/v2/logger"
)

func CheckMetricDef(id string, m *IndvMetric) error {
	_, err := GetMetricDefinition(id)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Debugf("adding %s to metric defs", id)
			_, err = NewFromMessage(m)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
