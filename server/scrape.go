// Copyright 2013 The Chihaya Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

package server

import (
	"bytes"

	cdb "github.com/pushrax/chihaya/database"
)

func writeScrapeInfo(torrent *cdb.Torrent, buf *bytes.Buffer) {
	buf.WriteRune('d')
	bencode("complete", buf)
	bencode(len(torrent.Seeders), buf)
	bencode("downloaded", buf)
	bencode(torrent.Snatched, buf)
	bencode("incomplete", buf)
	bencode(len(torrent.Leechers), buf)
	buf.WriteRune('e')
}

func scrape(params *queryParams, db *cdb.Database, buf *bytes.Buffer) {
	buf.WriteRune('d')
	bencode("files", buf)
	db.TorrentsMutex.RLock()
	if params.infoHashes != nil {
		for _, infoHash := range params.infoHashes {
			torrent, exists := db.Torrents[infoHash]
			if exists {
				bencode(infoHash, buf)
				writeScrapeInfo(torrent, buf)
			}
		}
	} else if infoHash, exists := params.get("info_hash"); exists {
		torrent, exists := db.Torrents[infoHash]
		if exists {
			bencode(infoHash, buf)
			writeScrapeInfo(torrent, buf)
		}
	}
	db.TorrentsMutex.RUnlock()
	buf.WriteRune('e')
}
