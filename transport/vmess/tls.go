package vmess

import (
	"context"
	"crypto/tls"
	"net"

	C "github.com/medasz/clash-kernel/constant"
)

type TLSConfig struct {
	Host           string
	SkipCertVerify bool
	NextProtos     []string
}

func StreamTLSConn(conn net.Conn, cfg *TLSConfig) (net.Conn, error) {
	tlsConfig := &tls.Config{
		ServerName:         cfg.Host,
		InsecureSkipVerify: cfg.SkipCertVerify,
		NextProtos:         cfg.NextProtos,
	}

	tlsConn := tls.Client(conn, tlsConfig)

	// fix tls handshake not timeout
	ctx, cancel := context.WithTimeout(context.Background(), C.DefaultTLSTimeout)
	defer cancel()
	err := tlsConn.HandshakeContext(ctx)
	return tlsConn, err
}
