package handlers

import (
    "net/http"
    "github.com/graphql-go/graphql"
    "github.com/labstack/echo/v4"
)

type GraphQLHandler struct {
    Schema *graphql.Schema
}

func (h *GraphQLHandler) ServeHTTP(c echo.Context) error {
    var params struct {
        Query string `json:"query"`
    }

    if err := c.Bind(&params); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    result := graphql.Do(graphql.Params{
        Schema:        *h.Schema,
        RequestString: params.Query,
    })

    if len(result.Errors) > 0 {
        return c.JSON(http.StatusBadRequest, result.Errors)
    }

    return c.JSON(http.StatusOK, result)
}
