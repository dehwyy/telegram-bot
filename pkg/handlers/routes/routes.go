package routes

// Routes with description

type Route string

func (r Route) String() string { return string(r) }

const (
	RouteDefaultStart Route = "/start"
)

const (
	RouteAdminStart Route = "/admin"
)

const (
	RouteOperatorStart Route = "/operator"
)

const (
	RouteProviderStart Route = "/provider"
)

type RouteWithDescription struct {
	Route       Route
	Description string
}

var RoutesOverview []RouteWithDescription = []RouteWithDescription{
	{Route: RouteDefaultStart, Description: "Открыть меню"},
	{Route: RouteAdminStart, Description: "Открыть панель администратора"},
	{Route: RouteOperatorStart, Description: "Открыть панель оператора"},
	{Route: RouteProviderStart, Description: "Открыть панель провайдера"},
}
