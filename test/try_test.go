package test

import (
	"github.com/farseer-go/fs/exception"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTry(t *testing.T) {
	exception.Try(func() {
		assert.Panics(t, func() {
			exception.ThrowExceptionf("%d", 123)
		})
		assert.Panics(t, func() {
			exception.ThrowRefuseExceptionf("%d", 123)
		})
		assert.Panics(t, func() {
			exception.ThrowWebExceptionf(502, "%d", 123)
		})
	}).CatchWebException(func(exp exception.WebException) {
		os.Exit(-1)
	}).CatchRefuseException(func(exp exception.RefuseException) {
		os.Exit(-1)
	}).CatchStringException(func(exp string) {
		os.Exit(-1)
	}).CatchException(func(exp any) {
		os.Exit(-1)
	})

	exception.Try(func() {
		exception.ThrowException("aaa")
	}).CatchWebException(func(exp exception.WebException) {
		os.Exit(-1)
	}).CatchRefuseException(func(exp exception.RefuseException) {
		os.Exit(-1)
	}).CatchStringException(func(exp string) {
		assert.Equal(t, "aaa", exp)
	}).CatchException(func(exp any) {
		os.Exit(-1)
	})

	exception.Try(func() {
		exception.ThrowRefuseException("aaa")
	}).CatchWebException(func(exp exception.WebException) {
		os.Exit(-1)
	}).CatchStringException(func(exp string) {
		os.Exit(-1)
	}).CatchRefuseException(func(exp exception.RefuseException) {
		assert.Equal(t, "aaa", exp.Message)
	}).CatchException(func(exp any) {
		os.Exit(-1)
	})

	exception.Try(func() {
		exception.ThrowWebException(502, "aaa")
	}).CatchStringException(func(exp string) {
		os.Exit(-1)
	}).CatchRefuseException(func(exp exception.RefuseException) {
		os.Exit(-1)
	}).CatchWebException(func(exp exception.WebException) {
		assert.Equal(t, "aaa", exp.Message)
		assert.Equal(t, 502, exp.StatusCode)
	}).CatchException(func(exp any) {
		os.Exit(-1)
	})

	try := exception.Try(func() {
		exception.ThrowRefuseException("aaa")
	})
	try.CatchRefuseException(func(exp exception.RefuseException) {
		panic("aaa")
	})
	try.CatchException(func(exp any) {
		assert.Equal(t, "aaa", exp)
		panic("aaa")
	})
	assert.Panics(t, func() {
		try.ThrowUnCatch()
	})
}
