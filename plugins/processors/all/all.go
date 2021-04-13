package all

import (
	// Blank imports for plugins to register themselves
	_ "github.com/influxdata/telegraf/plugins/processors/aws/ec2"
	_ "github.com/influxdata/telegraf/plugins/processors/clone"
	_ "github.com/influxdata/telegraf/plugins/processors/converter"
	_ "github.com/influxdata/telegraf/plugins/processors/date"
	_ "github.com/influxdata/telegraf/plugins/processors/dedup"
	_ "github.com/influxdata/telegraf/plugins/processors/defaults"
	_ "github.com/influxdata/telegraf/plugins/processors/enum"
	_ "github.com/influxdata/telegraf/plugins/processors/execd"
	_ "github.com/influxdata/telegraf/plugins/processors/filepath"
	_ "github.com/influxdata/telegraf/plugins/processors/ifname"
	_ "github.com/influxdata/telegraf/plugins/processors/override"
	_ "github.com/influxdata/telegraf/plugins/processors/parser"
	_ "github.com/influxdata/telegraf/plugins/processors/pivot"
	_ "github.com/influxdata/telegraf/plugins/processors/port_name"
	_ "github.com/influxdata/telegraf/plugins/processors/printer"
	_ "github.com/influxdata/telegraf/plugins/processors/regex"
	_ "github.com/influxdata/telegraf/plugins/processors/rename"
	_ "github.com/influxdata/telegraf/plugins/processors/reverse_dns"
	_ "github.com/influxdata/telegraf/plugins/processors/s2geo"
	_ "github.com/influxdata/telegraf/plugins/processors/starlark"
	_ "github.com/influxdata/telegraf/plugins/processors/strings"
	_ "github.com/influxdata/telegraf/plugins/processors/tag_limit"
	_ "github.com/influxdata/telegraf/plugins/processors/template"
	_ "github.com/influxdata/telegraf/plugins/processors/topk"
	_ "github.com/influxdata/telegraf/plugins/processors/unpivot"

	// erda related
	_ "github.com/influxdata/telegraf/plugins/processors/application_slow"
	_ "github.com/influxdata/telegraf/plugins/processors/clusterclock"
	_ "github.com/influxdata/telegraf/plugins/processors/compatibility"
	_ "github.com/influxdata/telegraf/plugins/processors/meta"
	_ "github.com/influxdata/telegraf/plugins/processors/point_space_to_underline"
	_ "github.com/influxdata/telegraf/plugins/processors/rate_limiter"
	_ "github.com/influxdata/telegraf/plugins/processors/tag_override"
)
