package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["rotatevec2"] = ROTATEVEC2 }

const ROTATEVEC2 = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_000039fd_00000000-9_rotatevec2.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/rotatevec2.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/usr/local/cuda/bin/../include/device_functions.h"

.visible .entry rotatevec2(
	.param .u64 rotatevec2_param_0,
	.param .u64 rotatevec2_param_1,
	.param .u64 rotatevec2_param_2,
	.param .u64 rotatevec2_param_3,
	.param .u64 rotatevec2_param_4,
	.param .u64 rotatevec2_param_5,
	.param .f32 rotatevec2_param_6,
	.param .u64 rotatevec2_param_7,
	.param .u64 rotatevec2_param_8,
	.param .u64 rotatevec2_param_9,
	.param .f32 rotatevec2_param_10,
	.param .u32 rotatevec2_param_11
)
{
	.reg .pred 	%p<3>;
	.reg .s32 	%r<21>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd10, [rotatevec2_param_0];
	ld.param.u64 	%rd11, [rotatevec2_param_1];
	ld.param.u64 	%rd12, [rotatevec2_param_2];
	ld.param.u64 	%rd13, [rotatevec2_param_3];
	ld.param.u64 	%rd14, [rotatevec2_param_4];
	ld.param.u64 	%rd15, [rotatevec2_param_5];
	ld.param.f32 	%f1, [rotatevec2_param_6];
	ld.param.u64 	%rd16, [rotatevec2_param_7];
	ld.param.u64 	%rd17, [rotatevec2_param_8];
	ld.param.u64 	%rd18, [rotatevec2_param_9];
	ld.param.f32 	%f2, [rotatevec2_param_10];
	ld.param.u32 	%r2, [rotatevec2_param_11];
	cvta.to.global.u64 	%rd1, %rd18;
	cvta.to.global.u64 	%rd2, %rd15;
	cvta.to.global.u64 	%rd3, %rd12;
	cvta.to.global.u64 	%rd4, %rd17;
	cvta.to.global.u64 	%rd5, %rd14;
	cvta.to.global.u64 	%rd6, %rd11;
	cvta.to.global.u64 	%rd7, %rd16;
	cvta.to.global.u64 	%rd8, %rd13;
	cvta.to.global.u64 	%rd9, %rd10;
	.loc 2 7 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 8 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 10 1
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd9, %rd19;
	add.s64 	%rd21, %rd8, %rd19;
	ld.global.f32 	%f3, [%rd21];
	ld.global.f32 	%f4, [%rd20];
	fma.rn.f32 	%f5, %f3, %f1, %f4;
	add.s64 	%rd22, %rd7, %rd19;
	ld.global.f32 	%f6, [%rd22];
	fma.rn.f32 	%f7, %f6, %f2, %f5;
	.loc 2 11 1
	add.s64 	%rd23, %rd6, %rd19;
	add.s64 	%rd24, %rd5, %rd19;
	ld.global.f32 	%f8, [%rd24];
	ld.global.f32 	%f9, [%rd23];
	fma.rn.f32 	%f10, %f8, %f1, %f9;
	add.s64 	%rd25, %rd4, %rd19;
	ld.global.f32 	%f11, [%rd25];
	fma.rn.f32 	%f12, %f11, %f2, %f10;
	.loc 2 12 1
	add.s64 	%rd26, %rd3, %rd19;
	add.s64 	%rd27, %rd2, %rd19;
	ld.global.f32 	%f13, [%rd27];
	ld.global.f32 	%f14, [%rd26];
	fma.rn.f32 	%f15, %f13, %f1, %f14;
	add.s64 	%rd28, %rd1, %rd19;
	ld.global.f32 	%f16, [%rd28];
	fma.rn.f32 	%f17, %f16, %f2, %f15;
	.loc 2 14 1
	mul.f32 	%f18, %f12, %f12;
	fma.rn.f32 	%f19, %f7, %f7, %f18;
	fma.rn.f32 	%f20, %f17, %f17, %f19;
	.loc 3 991 5
	sqrt.rn.f32 	%f21, %f20;
	.loc 2 15 1
	setp.eq.f32 	%p2, %f21, 0f00000000;
	selp.f32 	%f22, 0f3F800000, %f21, %p2;
	.loc 4 2399 3
	div.rn.f32 	%f23, %f7, %f22;
	.loc 2 17 1
	st.global.f32 	[%rd20], %f23;
	.loc 4 2399 3
	div.rn.f32 	%f24, %f12, %f22;
	.loc 2 18 1
	st.global.f32 	[%rd23], %f24;
	.loc 4 2399 3
	div.rn.f32 	%f25, %f17, %f22;
	.loc 2 19 1
	st.global.f32 	[%rd26], %f25;

BB0_2:
	.loc 2 21 2
	ret;
}


`
