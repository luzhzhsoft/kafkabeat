// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by cgo -godefs - DO NOT EDIT.
// cgo -godefs defs_arches_linux.go

package arch

import (
	"strconv"
)

const x32SyscallMask = 0x40000000

type AuditArch uint32

const (
	auditArchAARCH64     AuditArch = 0xc00000b7
	auditArchARM         AuditArch = 0x40000028
	auditArchARMEB       AuditArch = 0x28
	auditArchCRIS        AuditArch = 0x4000004c
	auditArchFRV         AuditArch = 0x5441
	auditArchI386        AuditArch = 0x40000003
	auditArchIA64        AuditArch = 0xc0000032
	auditArchM32R        AuditArch = 0x58
	auditArchM68K        AuditArch = 0x4
	auditArchMIPS        AuditArch = 0x8
	auditArchMIPS64      AuditArch = 0x80000008
	auditArchMIPS64N32   AuditArch = 0xa0000008
	auditArchMIPSEL      AuditArch = 0x40000008
	auditArchMIPSEL64    AuditArch = 0xc0000008
	auditArchMIPSEL64N32 AuditArch = 0xe0000008
	auditArchPARISC      AuditArch = 0xf
	auditArchPARISC64    AuditArch = 0x8000000f
	auditArchPPC         AuditArch = 0x14
	auditArchPPC64       AuditArch = 0x80000015
	auditArchPPC64LE     AuditArch = 0xc0000015
	auditArchS390        AuditArch = 0x16
	auditArchS390X       AuditArch = 0x80000016
	auditArchSH          AuditArch = 0x2a
	auditArchSH64        AuditArch = 0x8000002a
	auditArchSHEL        AuditArch = 0x4000002a
	auditArchSHEL64      AuditArch = 0xc000002a
	auditArchSPARC       AuditArch = 0x2
	auditArchSPARC64     AuditArch = 0x8000002b
	auditArchX86_64      AuditArch = 0xc000003e
)

var auditArchNames = map[AuditArch]string{
	auditArchAARCH64:     "aarch64",
	auditArchARM:         "arm",
	auditArchARMEB:       "armeb",
	auditArchCRIS:        "cris",
	auditArchFRV:         "frv",
	auditArchI386:        "i386",
	auditArchIA64:        "ia64",
	auditArchM32R:        "m32r",
	auditArchM68K:        "m68k",
	auditArchMIPS:        "mips",
	auditArchMIPS64:      "mips64",
	auditArchMIPS64N32:   "mips64n32",
	auditArchMIPSEL:      "mipsel",
	auditArchMIPSEL64:    "mipsel64",
	auditArchMIPSEL64N32: "mipsel64n32",
	auditArchPARISC:      "parisc",
	auditArchPARISC64:    "parisc64",
	auditArchPPC:         "ppc",
	auditArchPPC64:       "ppc64",
	auditArchPPC64LE:     "ppc64le",
	auditArchS390:        "s390",
	auditArchS390X:       "s390x",
	auditArchSH:          "sh",
	auditArchSH64:        "sh64",
	auditArchSHEL:        "shel",
	auditArchSHEL64:      "shel64",
	auditArchSPARC:       "sparc",
	auditArchSPARC64:     "sparc64",
	auditArchX86_64:      "x86_64",
}

func (a AuditArch) String() string {
	name, found := auditArchNames[a]
	if found {
		return name
	}

	return "unknown[" + strconv.Itoa(int(a)) + "]"
}
