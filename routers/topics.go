package routers

import (
	. "forum/controllers"
	. "forum/middlewares"
	. "forum/utils"
)

var TopicRoutes = Routes{
	Route{
		Method:  "GET",
		Path:    "",
		Handler: TopicList,
	},
	Route{
		Method:      "GET",
		Path:        "/:id",
		MiddleWares: MiddleWares{IDValidator},
		Handler:     TopicDetail,
	},
	Route{
		Method:  "POST",
		Path:    "",
		Handler: CreateTopic,
	},
	Route{
		Method:      "PUT",
		Path:        "/:id",
		MiddleWares: MiddleWares{IDValidator},
		Handler:     EditTopic,
	},
	Route{
		Method:      "DELETE",
		Path:        "/:id",
		MiddleWares: MiddleWares{IDValidator},
		Handler:     DeleteTopic,
	},
}
