// Copyright 2013 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package websocket implements the WebSocket protocol defined in RFC 6455.
//
// Overview
//
// The Conn type represents a WebSocket connection.
//
// A server application calls the Upgrade function to get a pointer to a Conn:
//
//  func handler(w http.ResponseWriter, r *http.Request) {
//      conn, err := websocket.Upgrade(w, r.Header, nil, 1024, 1024)
//      if _, ok := err.(websocket.HandshakeError); ok {
//          http.Error(w, "Not a websocket handshake", 400)
//          return
//      } else if err != nil {
//          log.Println(err)
//          return
//      }
//      ... Use conn to send and receive messages.
//  }
//
// WebSocket messages are represented by the io.Reader interface when receiving
// a message and by the io.WriteCloser interface when sending a message. An
// application receives a message by calling the Conn.NextReader method and
// reading the returned io.Reader to EOF. An application sends a message by
// calling the Conn.NextWriter method and writing the message to the returned
// io.WriteCloser. The application terminates the message by closing the
// io.WriteCloser.
//
// The following example shows how to use the connection NextReader and
// NextWriter method to echo messages:
//
//  for {
//      mt, r, err := conn.NextReader()
//      if err != nil {
//          return
//      }
//      w, err := conn.NextWriter(mt)
//      if err != nil {
//          return err
//      }
//      if _, err := io.Copy(w, r); err != nil {
//          return err
//      }
//      if err := w.Close(); err != nil {
//          return err
//      }
//  }
//
// The connection ReadMessage and WriteMessage methods are helpers for reading
// or writing an entire message in one method call. The following example shows
// how to echo messages using these connection helper methods:
//
//  for {
//      mt, p, err := conn.ReadMessage()
//      if err != nil {
//          return
//      }
//      if _, err := conn.WriteMessaage(mt, p); err != nil {
//          return err
//      }
//  }
//
// Data Message Types
//
// The WebSocket protocol distinguishes between text and binary data messages.
// Text messages are interpreted as UTF-8 encoded text. The interpretation of
// binary messages is left to the application.
//
// This package uses the same types and methods to work with both types of data
// messages. When sending or receiving a text message, it is the application's
// responsibility to ensure that the messages is valid UTF-8 encoded text.
//
// Concurrency
//
// A Conn supports a single concurrent caller to the write methods (NextWriter,
// SetWriteDeadline, WriteMessage) and a single concurrent caller to the read
// methods (NextReader, SetReadDeadline, ReadMessage). The Close and
// WriteControl methods can be called concurrently with all other methods.
package websocket
