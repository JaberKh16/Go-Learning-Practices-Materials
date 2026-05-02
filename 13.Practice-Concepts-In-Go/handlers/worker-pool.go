package handlers

type Sitemap struct {
	Loc     string
	LastMod string
}

type Job struct {
	URL string
}

type Result struct {
	URL    string
	Status string
	Err    error
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		resp, err := http.Get(job.URL)
		if err != nil {
			results <- Result{URL: job.URL, Status: "FAILED", Err: err}
			continue
		}
		resp.Body.Close()

		results <- Result{
			URL:    job.URL,
			Status: "OK",
			Err:    nil,
		}
	}
}

func FetchAllSitemapsWorkerPool(sitemaps []Sitemap, workerCount int) {
	jobs := make(chan Job, len(sitemaps))
	results := make(chan Result, len(sitemaps))

	// start workers
	for i := 0; i < workerCount; i++ {
		go worker(i, jobs, results)
	}

	// send jobs
	for _, s := range sitemaps {
		jobs <- Job{URL: s.Loc}
	}
	close(jobs)

	// collect results
	for i := 0; i < len(sitemaps); i++ {
		res := <-results
		if res.Err != nil {
			fmt.Println("❌ Failed:", res.URL, res.Err)
		} else {
			fmt.Println("✅ Fetched:", res.URL)
		}
	}
}

func PerformRequestWorkerPool() {
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var siteMapIndex SiteMapIndex
	err = xml.Unmarshal(body, &siteMapIndex)
	if err != nil {
		panic(err)
	}

	FetchAllSitemapsWorkerPool(siteMapIndex.Sitemaps, 5)
}

func WorkingWithXMLData(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var siteMapIndex SiteMapIndex
	xml.Unmarshal(body, &siteMapIndex)

	FetchAllSitemapsWorkerPool(siteMapIndex.Sitemaps, 5)

	fmt.Fprintln(w, "Worker pool execution completed")
}