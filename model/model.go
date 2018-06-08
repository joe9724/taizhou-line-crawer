package model

type RouteModel struct{
	RouteID int64 `json:"routeId"`
	RouteName string `json:"routeName"`
	IsHaveSubRouteCombine string `json:"isHaveSubRouteCombine"`
	RouteNameExt string `json:"RouteNameExt"`
}

type RouteData struct{
	TimeStamp string `json:"timestamp"`
	RouteList []RouteModel `json:"routeList"`
}
