package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bufio"
	"github.com/rgbkrk/libvirt-go"
	"html"
	"strings"
)

type virtlauncher struct {
	domainXML  string
	domainUUID string
	connURI    string
	user       string
	pass       string
}

func main() {
	xmlPath := flag.String("domain-path", "/var/run/virt-launcher/dom.xml", "Where to look for the domain xml.")
	downwardAPIPath := flag.String("downward-api-path", "", "Load domain from this downward API file")
	conUri := flag.String("libvirt-uri", "qemu:///system", "Libvirt connection string.")
	user := flag.String("user", "vdsm@ovirt", "Libvirt user")
	pass := flag.String("pass", "shibboleth", "Libvirt password")
	receiveOnly := flag.Bool("receive-only", false, "Do not create the domain")
	flag.Parse()

	launcher := virtlauncher{
		connURI: *conUri,
		user:    *user,
		pass:    *pass,
	}

	if !*receiveOnly {
		launcher.ReadDomainXML(*xmlPath, *downwardAPIPath)
		launcher.CreateDomain()
	}

	waitUntilSignal()
}

func (vl *virtlauncher) CreateDomain() {
	conn, err := libvirt.NewVirConnectionWithAuth(vl.connURI, vl.user, vl.pass)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to libvirt using %s: %s", vl.connURI, err))
	}
	defer func() {
		if _, closeErr := conn.CloseConnection(); closeErr != nil {
			log.Fatalf("CloseConnection() failed: %s", closeErr)
		}
		log.Print("Connection closed")
	}()

	log.Print("Libvirt connection established")

	// Launch VM
	_, err = conn.DomainCreateXML(vl.domainXML, 0)
	if err != nil {
		panic(fmt.Sprintf("Could not create the libvirt domain: %s", err))
	}

	log.Print("Domain started")
}

func waitUntilSignal() {
	// Wait for termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	log.Printf("Waiting forever...")
	s := <-c
	log.Print("Got signal: ", s)
}

func (vl *virtlauncher) ReadDomainXML(xmlPath string, downwardAPIPath string) {
	if downwardAPIPath == "" {
		log.Print("Loading Domain from XML file")
		rawXML, err := ioutil.ReadFile(xmlPath)
		if err != nil {
			log.Fatal(err)
		}
		vl.domainXML = string(rawXML)
	} else {
		log.Print("Loading Domain from downward API file")
		f, err := os.Open(downwardAPIPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, `domainXML="`) {
				vl.domainXML = DecodeDomainXML(strings.Trim(strings.TrimPrefix(line, "domainXML="), `"`))
			}
		}

	}

	if vl.domainXML == "" {
		panic("Could not load domain XML. The resulting XML is empty")
	}
	log.Print("Domain description loaded.")
}

func DecodeDomainXML(domainXML string) string {
	decodedXML := html.UnescapeString(string(domainXML))
	decodedXML = strings.Replace(decodedXML, "\\\\", "\\", -1)
	decodedXML = strings.Replace(decodedXML, "\\n", "\n", -1)
	return decodedXML
}