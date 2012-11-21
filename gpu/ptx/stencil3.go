package ptx

const STENCIL3 = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00002274_00000000-9_stencil3.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/stencil3.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"

.visible .entry stencil3(
	.param .u64 stencil3_param_0,
	.param .u64 stencil3_param_1,
	.param .f32 stencil3_param_2,
	.param .f32 stencil3_param_3,
	.param .f32 stencil3_param_4,
	.param .f32 stencil3_param_5,
	.param .f32 stencil3_param_6,
	.param .f32 stencil3_param_7,
	.param .f32 stencil3_param_8,
	.param .u32 stencil3_param_9,
	.param .u32 stencil3_param_10,
	.param .u32 stencil3_param_11,
	.param .u32 stencil3_param_12,
	.param .u32 stencil3_param_13,
	.param .u32 stencil3_param_14
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<73>;
	.reg .f32 	%f<24>;
	.reg .s64 	%rd<20>;


	ld.param.u64 	%rd3, [stencil3_param_0];
	ld.param.u64 	%rd4, [stencil3_param_1];
	ld.param.f32 	%f1, [stencil3_param_2];
	ld.param.f32 	%f2, [stencil3_param_3];
	ld.param.f32 	%f3, [stencil3_param_4];
	ld.param.f32 	%f4, [stencil3_param_5];
	ld.param.f32 	%f5, [stencil3_param_6];
	ld.param.f32 	%f6, [stencil3_param_7];
	ld.param.f32 	%f7, [stencil3_param_8];
	ld.param.u32 	%r28, [stencil3_param_12];
	ld.param.u32 	%r29, [stencil3_param_13];
	ld.param.u32 	%r30, [stencil3_param_14];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 25 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 26 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 28 1
	setp.lt.s32 	%p1, %r8, %r30;
	setp.lt.s32 	%p2, %r4, %r29;
	and.pred  	%p3, %p2, %p1;
	.loc 2 32 1
	setp.gt.s32 	%p4, %r28, 0;
	.loc 2 28 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 36 1
	add.s32 	%r9, %r28, -1;
	.loc 2 38 1
	add.s32 	%r32, %r4, 1;
	mov.u32 	%r31, 0;
	.loc 3 238 5
	max.s32 	%r33, %r32, %r31;
	.loc 2 38 1
	add.s32 	%r34, %r29, -1;
	.loc 3 210 5
	min.s32 	%r35, %r33, %r34;
	.loc 2 39 1
	add.s32 	%r36, %r4, -1;
	.loc 3 238 5
	max.s32 	%r37, %r36, %r31;
	.loc 3 210 5
	min.s32 	%r38, %r37, %r34;
	.loc 2 40 1
	add.s32 	%r39, %r8, 1;
	.loc 3 238 5
	max.s32 	%r40, %r39, %r31;
	.loc 2 40 1
	add.s32 	%r41, %r30, -1;
	.loc 3 210 5
	min.s32 	%r42, %r40, %r41;
	.loc 2 41 1
	add.s32 	%r43, %r8, -1;
	.loc 3 238 5
	max.s32 	%r44, %r43, %r31;
	.loc 3 210 5
	min.s32 	%r45, %r44, %r41;
	.loc 2 32 1
	mad.lo.s32 	%r71, %r38, %r30, %r8;
	mul.lo.s32 	%r11, %r30, %r29;
	mad.lo.s32 	%r70, %r35, %r30, %r8;
	mad.lo.s32 	%r69, %r30, %r4, %r45;
	mad.lo.s32 	%r68, %r30, %r4, %r42;
	mad.lo.s32 	%r67, %r30, %r4, %r8;
	mov.u32 	%r72, %r31;

BB0_2:
	.loc 2 34 1
	mov.u32 	%r21, %r72;
	mul.wide.s32 	%rd5, %r67, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.f32 	%f8, [%rd6];
	.loc 2 36 1
	add.s32 	%r22, %r21, 1;
	.loc 3 238 5
	max.s32 	%r50, %r22, %r31;
	.loc 3 210 5
	min.s32 	%r51, %r50, %r9;
	.loc 2 36 1
	mad.lo.s32 	%r52, %r51, %r29, %r4;
	mad.lo.s32 	%r53, %r52, %r30, %r8;
	mul.wide.s32 	%rd7, %r53, 4;
	add.s64 	%rd8, %rd2, %rd7;
	ld.global.f32 	%f9, [%rd8];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f1, %f10;
	.loc 2 37 1
	add.s32 	%r55, %r21, -1;
	.loc 3 238 5
	max.s32 	%r56, %r55, %r31;
	.loc 3 210 5
	min.s32 	%r57, %r56, %r9;
	.loc 2 37 1
	mad.lo.s32 	%r58, %r57, %r29, %r4;
	mad.lo.s32 	%r59, %r58, %r30, %r8;
	mul.wide.s32 	%rd9, %r59, 4;
	add.s64 	%rd10, %rd2, %rd9;
	ld.global.f32 	%f12, [%rd10];
	fma.rn.f32 	%f13, %f12, %f7, %f11;
	.loc 2 38 1
	mul.wide.s32 	%rd11, %r70, 4;
	add.s64 	%rd12, %rd2, %rd11;
	ld.global.f32 	%f14, [%rd12];
	fma.rn.f32 	%f15, %f14, %f4, %f13;
	.loc 2 39 1
	mul.wide.s32 	%rd13, %r71, 4;
	add.s64 	%rd14, %rd2, %rd13;
	ld.global.f32 	%f16, [%rd14];
	fma.rn.f32 	%f17, %f16, %f5, %f15;
	.loc 2 40 1
	mul.wide.s32 	%rd15, %r68, 4;
	add.s64 	%rd16, %rd2, %rd15;
	ld.global.f32 	%f18, [%rd16];
	fma.rn.f32 	%f19, %f18, %f3, %f17;
	.loc 2 41 1
	mul.wide.s32 	%rd17, %r69, 4;
	add.s64 	%rd18, %rd2, %rd17;
	ld.global.f32 	%f20, [%rd18];
	fma.rn.f32 	%f21, %f20, %f2, %f19;
	.loc 2 43 1
	add.s64 	%rd19, %rd1, %rd5;
	ld.global.f32 	%f22, [%rd19];
	add.f32 	%f23, %f22, %f21;
	st.global.f32 	[%rd19], %f23;
	.loc 2 32 1
	add.s32 	%r71, %r71, %r11;
	add.s32 	%r70, %r70, %r11;
	add.s32 	%r69, %r69, %r11;
	add.s32 	%r68, %r68, %r11;
	add.s32 	%r67, %r67, %r11;
	setp.lt.s32 	%p6, %r22, %r28;
	mov.u32 	%r72, %r22;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 45 2
	ret;
}


`
