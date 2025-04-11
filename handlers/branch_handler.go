package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// BranchHandler handles HTTP requests for branches
type BranchHandler struct {
	branchFirebase *models.BranchFirebase
}

// NewBranchHandler creates a new BranchHandler instance
func NewBranchHandler(branchFirebase *models.BranchFirebase) *BranchHandler {
	return &BranchHandler{
		branchFirebase: branchFirebase,
	}
}

// List handles GET requests to list all branches
func (h *BranchHandler) List(c *gin.Context) {
	branches, err := h.branchFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branches)
}

// Get handles GET requests to get a specific branch
func (h *BranchHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required"})
		return
	}

	branch, err := h.branchFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branch)
}

// Create handles POST requests to create a new branch
func (h *BranchHandler) Create(c *gin.Context) {
	var branch models.BranchFirebaseModel
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.branchFirebase.Create(c.Request.Context(), &branch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	branch.ID = id
	c.JSON(http.StatusCreated, branch)
}

// Update handles PUT requests to update a branch
func (h *BranchHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required"})
		return
	}

	var branch models.BranchFirebaseModel
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch.ID = id
	if err := h.branchFirebase.Update(c.Request.Context(), id, &branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branch)
}

// Delete handles DELETE requests to delete a branch
func (h *BranchHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID is required"})
		return
	}

	if err := h.branchFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch deleted successfully"})
}
