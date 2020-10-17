package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "Website Lookup CLI"
    app.Usage = "Web Domain Tools"
    app.Description = "Standard Web Domain Lookup Tools"
    app.Version = "1.0.0"

    // We'll be using the same flag for all our commands
    // so we'll define it up here
    myFlags := []cli.Flag{
        cli.StringFlag{
            Name:  "url",
            Value: "google.com",
        },
    }

    // we create our commands
    app.Commands = []cli.Command{
        // NS
        {
            Name:  "ns",
            Usage: "./domain_tool ns --url URL - Looks Up the NameServers",
            Flags: myFlags,
            // the action, or code that will be executed when
            // we execute our `ns` command
            Action: func(c *cli.Context) error {
                // a simple lookup function
                ns, err := net.LookupNS(c.String("url"))
                if err != nil {
                    return err
                }
                // we log the results to our console
                // using a trusty fmt.Println statement
                for i := 0; i < len(ns); i++ {
                    fmt.Println(ns[i].Host)
                }
                return nil
            },
        },
        
        // IP
        {
            Name:  "ip",
            Usage: "./domain_tool ip --url URL - Looks up the IP addresses",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
            ip, err := net.LookupIP(c.String("url"))
            if err != nil {
               fmt.Println(err)
            }
            for i := 0; i < len(ip); i++ {
               fmt.Println(ip[i])
            }
            return nil
            },
        },
        
        {
    	   Name:  "cname",
    	   Usage: "./domain_tool cname --url URL - Looks up the CNAME",
    	   Flags: myFlags,
    	   Action: func(c *cli.Context) error {
        	cname, err := net.LookupCNAME(c.String("url"))
        	if err != nil {
            	   fmt.Println(err)
        	}
        	fmt.Println(cname)
        	return nil
    	    },
	},
	
	// Host
	{
    	   Name:  "domain",
    	   Usage: "./domain_tool domain --url URL - Looks up a domain name",
    	   Flags: myFlags,
    	   Action: func(c *cli.Context) error {
        	cname, err := net.LookupHost(c.String("url"))
        	if err != nil {
            	   fmt.Println(err)
        	}
        	fmt.Println(cname)
        	return nil
    	    },
	},
	
	// Reverse
	{
    	   Name:  "reverse",
    	   Usage: "./domain_tool reverse --url URL - Performs a reverse lookup",
    	   Flags: myFlags,
    	   Action: func(c *cli.Context) error {
        	cname, err := net.LookupHost(c.String("url"))
        	if err != nil {
            	   fmt.Println(err)
        	}
        	fmt.Println(cname)
        	return nil
    	    },
	},
	
	// Text
	{
    	   Name:  "txt",
    	   Usage: "./domain_tool txt --url URL - Looks up txt records",
    	   Flags: myFlags,
    	   Action: func(c *cli.Context) error {
        	cname, err := net.LookupTXT(c.String("url"))
        	if err != nil {
            	   fmt.Println(err)
        	}
        	fmt.Println(cname)
        	return nil
    	    },
	},
	
	{
    Name:  "mx",
    Usage: "./domain_tool mx --url URL - Looks up the MX records",
    Flags: myFlags,
    Action: func(c *cli.Context) error {
        mx, err := net.LookupMX(c.String("url"))
        if err != nil {
            fmt.Println(err)
        }
        for i := 0; i < len(mx); i++ {
            fmt.Println(mx[i].Host, mx[i].Pref)
        }
        return nil
    },
},
    }

    // start our application
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
