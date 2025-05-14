package main

import (
	"log"
	"net/http"
	"github.com/Scrimzay/spacesite/nasaapi"
	"github.com/Scrimzay/spacesite/spacexapi"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/**/*.html")
	r.Static("/static", "./static")

	r.GET("/", indexHandler)
	r.GET("/spacex", spaceXIndexHandler)
	r.GET("/nasa", nasaIndexHandler)
    r.GET("/spacex/rockets", spaceXRocketHandler)
    r.GET("/spacex/rockets/info", generalSpaceXRocketInfo)
    r.POST("/spacex/rockets/search", searchRocketByID)
    r.GET("/spacex/rockets/:id/info", specificSpaceXRocketInfo)
    r.GET("/spacex/starlink", spaceXStarlinkHandler)
    r.GET("/spacex/starlink/info", generalSpaceXStarlinkInfo)
    r.POST("/spacex/starlink/search", searchStarlinkByID)
    r.GET("/spacex/starlink/:id/info", specificSpaceXStarlinkInfo)
    r.GET("/spacex/roadster", spaceXRoadsterHandler)
    r.GET("/spacex/roadster/info", generalSpaceXRoadsterInfo)
    r.GET("/spacex/company", spaceXCompanyHandler)
    r.GET("/spacex/company/info", generalSpaceXCompanyInfo)
    r.GET("/spacex/crew", spaceXCrewHandler)
    r.GET("/spacex/crew/info", generalSpaceXCrewInfo)
    r.GET("/spacex/dragons", spaceXDragonsHandler)
    r.GET("/spacex/dragons/info", generalSpaceXDragonsInfo)
    r.GET("/nasa/tle", nasaTLEHandler)
    r.GET("/nasa/tle/info", generalNASATLEInformation)

	err := r.Run(":5600")
	if err != nil {
		log.Fatalf("Could not start: %v", err)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func spaceXIndexHandler(c *gin.Context) {
	c.HTML(200, "spacexindex.html", nil)
}

func nasaIndexHandler(c *gin.Context) {
	c.HTML(200, "nasaindex.html", nil)
}

func spaceXRocketHandler(c *gin.Context) {
    c.HTML(200, "spacexrocketindex.html", nil)
}

func spaceXStarlinkHandler(c *gin.Context) {
    c.HTML(200, "spacexstarlinkindex.html", nil)
}

func spaceXRoadsterHandler(c *gin.Context) {
    c.HTML(200, "spacexroadsterindex.html", nil)
}

func spaceXCompanyHandler(c *gin.Context) {
    c.HTML(200, "spacexcompanyindex.html", nil)
}

func spaceXCrewHandler(c *gin.Context) {
    c.HTML(200, "spacexcrewindex.html", nil)
}

func spaceXDragonsHandler(c *gin.Context) {
    c.HTML(200, "spacexdragonsindex.html", nil)
}

func nasaTLEHandler(c *gin.Context) {
    c.HTML(200, "nasatleindex.html", nil)
}

func generalSpaceXRocketInfo(c *gin.Context) {
    rockets, err := spacexapi.CallGeneralSpaceXRocketsInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX rockets: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexrocketinfo.html", gin.H{
            "Error": "Failed to fetch rocket info",
        })
    }

    c.HTML(200, "generalspacexrocketinfo.html", gin.H{
        "Rockets": rockets,
    })
}

func specificSpaceXRocketInfo(c *gin.Context) {
    id := c.Param("id")
    rocket, err := spacexapi.CallSpecificSpaceXRocketsInformation(id)
    if err != nil {
        c.HTML(http.StatusNotFound, "specificspacexrocketinfo.html", gin.H{
            "Error": "Rocket not found",
        })
        return
    }

    c.HTML(http.StatusOK, "specificspacexrocketinfo.html", gin.H{
        "Rocket": rocket,
    })
}

func searchRocketByID(c *gin.Context) {
    id := c.PostForm("id")
    id = strings.TrimSpace(id)
    if id == "" {
        log.Printf("Empty rocket ID in form")
        c.HTML(http.StatusBadRequest, "spacexrocketindex.html", gin.H{
            "Error": "Rocket ID is required",
        })
        return
    }

    // Validate ID by attempting to fetch the rocket
    rocket, err := spacexapi.CallSpecificSpaceXRocketsInformation(id)
    if err != nil {
        log.Printf("Error validating rocket ID %s: %v", id, err)
        c.HTML(http.StatusNotFound, "spacexrocketindex.html", gin.H{
            "Error": "Rocket not found for ID: " + id,
        })
        return
    }

    redirectURL := "/spacex/rockets/" + rocket.ID + "/info"
    log.Printf("Redirecting to %s for rocket ID %s", redirectURL, id)
    c.Redirect(http.StatusFound, redirectURL)
}

func generalSpaceXStarlinkInfo(c *gin.Context) {
    satellites, err := spacexapi.CallGeneralSpaceXStarlinkInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX starlink: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexstarlinkinfo.html", gin.H{
            "Error": "Failed to fetch starlink info",
        })
    }

    c.HTML(200, "generalspacexstarlinkinfo.html", gin.H{
        "Satellites": satellites,
    })
}

func specificSpaceXStarlinkInfo(c *gin.Context) {
    id := c.Param("id")
    satellite, err := spacexapi.CallSpecificSpaceXStarlinkInformation(id)
    if err != nil {
        c.HTML(http.StatusNotFound, "specificspacexstarlinkinfo.html", gin.H{
            "Error": "Satellite not found",
        })
        return
    }

    c.HTML(http.StatusOK, "specificspacexstarlinkinfo.html", gin.H{
        "Satellite": satellite,
    })
}

func searchStarlinkByID(c *gin.Context) {
    id := c.PostForm("id")
    id = strings.TrimSpace(id)
    if id == "" {
        log.Printf("Empty satellite ID in form")
        c.HTML(http.StatusBadRequest, "spacexstarlinkindex.html", gin.H{
            "Error": "Satellite ID is required",
        })
        return
    }

    // Validate ID by attempting to fetch the rocket
    satellite, err := spacexapi.CallSpecificSpaceXStarlinkInformation(id)
    if err != nil {
        log.Printf("Error validating satellite ID %s: %v", id, err)
        c.HTML(http.StatusNotFound, "spacexsatelliteindex.html", gin.H{
            "Error": "satellite not found for ID: " + id,
        })
        return
    }

    redirectURL := "/spacex/starlink/" + satellite.ID + "/info"
    c.Redirect(http.StatusFound, redirectURL)
}

func generalSpaceXRoadsterInfo(c *gin.Context) {
    roadster, err := spacexapi.CallGeneralSpaceXRoadsterInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX roadster: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexroadsterinfo.html", gin.H{
            "Error": "Failed to fetch roadster info",
        })
    }

    c.HTML(200, "generalspacexroadsterinfo.html", gin.H{
        "Roadster": roadster,
    })
}

func generalSpaceXCompanyInfo(c *gin.Context) {
    company, err := spacexapi.CallGeneralSpaceXCompanyInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX company: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexcompanyinfo.html", gin.H{
            "Error": "Failed to fetch company info",
        })
    }

    c.HTML(200, "generalspacexcompanyinfo.html", gin.H{
        "Company": company,
    })
}

func generalSpaceXCrewInfo(c *gin.Context) {
    crew, err := spacexapi.CallGeneralSpaceXCrewInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX crew: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexcrewinfo.html", gin.H{
            "Error": "Failed to fetch crew info",
        })
    }

    c.HTML(200, "generalspacexcrewinfo.html", gin.H{
        "Crew": crew,
    })
}

func generalSpaceXDragonsInfo(c *gin.Context) {
    dragons, err := spacexapi.CallGeneralSpaceXDragonsInformation()
    if err != nil {
        log.Printf("Error fetching SpaceX dragons: %v", err)
        c.HTML(http.StatusInternalServerError, "generalspacexdragonsinfo.html", gin.H{
            "Error": "Failed to fetch dragons info",
        })
    }

    c.HTML(200, "generalspacexdragonsinfo.html", gin.H{
        "Dragons": dragons,
    })
}

func generalNASATLEInformation(c *gin.Context) {
    tles, err := nasaapi.CallGeneralNASATLEInformation()
    if err != nil {
        log.Printf("Error fetching NASA TLEs: %v", err)
        c.HTML(http.StatusInternalServerError, "generalnasatleinfo.html", gin.H{
            "Error": "Failed to fetch TLEs info",
        })
    }

    c.HTML(200, "generalnasatleinfo.html", gin.H{
        "TLEs": tles,
    })
}