package routers

import (
	. "forum/controllers"
)

var TopicRoutes = routes{
	route{
		method:  "GET",
		path:    "",
		handler: TopicList,
	},
	route{
		method:  "GET",
		path:    "/:id",
		handler: TopicDetail,
	},
	route{
		method:  "POST",
		path:    "",
		handler: CreateTopic,
	},
	route{
		method:  "PUT",
		path:    "/:id",
		handler: EditTopic,
	},
	route{
		method:  "DELETE",
		path:    "/:id",
		handler: DeleteTopic,
	},
}
