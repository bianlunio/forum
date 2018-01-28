package routers

import (
	. "forum/controllers"
)

var TopicRoutes = routes{
	route{
		Method:      "GET",
		Path:        "",
		Handler:     TopicList,
	},
	//Route{
	//	Method:      "GET",
	//	Path:        "/:id",
	//	MiddleWares: MiddleWares{IDValidator},
	//	Handler:     TopicDetail,
	//},
	route{
		Method:  "POST",
		Path:    "",
		Handler: CreateTopic,
	},
	//Route{
	//	Method:      "PUT",
	//	Path:        "/:id",
	//	MiddleWares: MiddleWares{IDValidator},
	//	Handler:     EditTopic,
	//},
	//Route{
	//	Method:      "DELETE",
	//	Path:        "/:id",
	//	MiddleWares: MiddleWares{IDValidator},
	//	Handler:     DeleteTopic,
	//},
}
