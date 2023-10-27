package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"project/auth"
	"project/middlewares"
	"project/models"
)

func (h *handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	claims, ok := ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var newInv models.CreateCompany
	err := json.NewDecoder(c.Request.Body).Decode(&newInv)
	if err != nil {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	validate := validator.New()
	err = validate.Struct(newInv)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"msg": "please provide Item Name and Quantity"})
		return
	}
	uid, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.
			StatusInternalServerError)})
		return
	}
	inv, err := h.s.CreateCompany(ctx, newInv, uint(uid))
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Inventory creation failed"})
		return
	}

	c.JSON(http.StatusOK, inv)

}

func (h *handler) ViewCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	claims, ok := ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	company, err := h.s.ViewCompany(ctx, claims.Subject)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing company"})
		return
	}
	c.JSON(http.StatusOK, company)
}

// func (h *handler) ViewCompany(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
// 	if !ok {
// 		log.Error().Msg("traceId missing from context")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
// 		return
// 	}

// 	claims, ok := ctx.Value(auth.Key).(jwt.RegisteredClaims)
// 	if !ok {
// 		log.Error().Str("Trace Id", traceId).Msg("login first")
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
// 		return
// 	}
// 	company, err := h.s.ViewCompany(ctx, claims.Subject)

// 	if err != nil {
// 		log.Error().Err(err).Str("Trace Id", traceId)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing company"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, company)
// }

func (h *handler) viewSingleCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceid", traceid).Msg("traceid is not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}
	id, err := strconv.ParseInt(c.Param("ComapanyId"), 10, 64)
	if err != nil {
		log.Error().Msg("id is not converting")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}
	company, err := h.s.Getcompany(id)
	if err != nil {
		log.Error().Err(err).Msg("error in fecting the company by id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing company"})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *handler) postJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceid", traceid).Msg("traceid is not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	id, err := strconv.ParseUint(c.Param("company_id"), 10, 64)
	if err != nil {
		log.Error().Str("traceid", traceid).Msg("traceid is not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}
	var jobCreation models.CreateJob
	body := c.Request.Body
	err = json.NewDecoder(body).Decode(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}
	us, err := h.s.JobCreation(ctx, jobCreation, id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceid).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}
