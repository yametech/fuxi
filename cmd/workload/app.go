package main

import "github.com/gin-gonic/gin"

func ListCharts(g *gin.Context) { workloadsAPI.ListCharts(g) }

func GetCharts(g *gin.Context) { workloadsAPI.GetCharts(g) }

func GetChartValues(g *gin.Context) { workloadsAPI.GetChartValues(g) }

func InstallChart(g *gin.Context) { workloadsAPI.InstallChart(g) }

func ListRelease(g *gin.Context) { workloadsAPI.ListRelease(g) }

func FindReleasesByNamespace(g *gin.Context) { workloadsAPI.FindReleasesByNamespce(g) }

func FindReleaseByNamespace(g *gin.Context) { workloadsAPI.FindReleaseByNamespce(g) }

func FindReleaseValueByName(g *gin.Context) { workloadsAPI.FindReleaseValueByName(g) }

func DeleteRelease(g *gin.Context) { workloadsAPI.DeleteRelease(g) }

func UpgradeRelease(g *gin.Context) { workloadsAPI.UpgradeRelease(g) }

func RollbackRelease(g *gin.Context) { workloadsAPI.RollbackRelease(g) }

func HistoryRelease(g *gin.Context) { workloadsAPI.HistoryRelease(g) }
