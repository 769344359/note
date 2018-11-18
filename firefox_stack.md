```
Thread 1 "firefox" hit Breakpoint 2, mozilla::dom::Element_Binding::querySelectorAll (cx=0x7f39dea24000, obj=..., self=0x7f39dd15e4c0, 
    args=...) at /home/dinosaur/Downloads/fire/dom/bindings/ElementBinding.cpp:3780
3780	{
(gdb) bt
#0  mozilla::dom::Element_Binding::querySelectorAll (cx=0x7f39dea24000, obj=..., self=0x7f39dd15e4c0, args=...)
    at /home/dinosaur/Downloads/fire/dom/bindings/ElementBinding.cpp:3780
#1  0x00007f39e5648e71 in mozilla::dom::binding_detail::GenericMethod<mozilla::dom::binding_detail::NormalThisPolicy, mozilla::dom::binding_detail::ThrowExceptions> (cx=<optimized out>, argc=<optimized out>, vp=<optimized out>)
    at /home/dinosaur/Downloads/firefox-64.0/dom/bindings/BindingUtils.cpp:3315
#2  0x00007f39e797832b in CallJSNative (cx=0x7f39dea24000, 
    native=0x7f39e5648ceb <mozilla::dom::binding_detail::GenericMethod<mozilla::dom::binding_detail::NormalThisPolicy, mozilla::dom::binding_detail::ThrowExceptions>(JSContext*, unsigned int, JS::Value*)>, args=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:468
#3  0x00007f39e798b13f in js::InternalCallOrConstruct (cx=<optimized out>, cx@entry=0x7f39dea24000, args=..., 
    construct=construct@entry=js::NO_CONSTRUCT) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:560
#4  0x00007f39e798b78d in InternalCall (cx=cx@entry=0x7f39dea24000, args=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:614
#5  0x00007f39e798b900 in js::Call (cx=cx@entry=0x7f39dea24000, fval=..., fval@entry=..., thisv=..., args=..., rval=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:633
#6  0x00007f39e7485945 in js::ForwardingProxyHandler::call (this=<optimized out>, cx=0x7f39dea24000, proxy=..., args=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/proxy/Wrapper.cpp:178
#7  0x00007f39e74723e3 in js::CrossCompartmentWrapper::call (this=0x7f39eac691a0 <js::CrossCompartmentWrapper::singleton>, 
    cx=<optimized out>, wrapper=..., args=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/proxy/CrossCompartmentWrapper.cpp:355
#8  0x00007f39e747a2a2 in js::Proxy::call (cx=0x7f39dea24000, proxy=proxy@entry=..., args=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/proxy/Proxy.cpp:560
#9  0x00007f39e798b5e6 in js::InternalCallOrConstruct (cx=<optimized out>, cx@entry=0x7f39dea24000, args=..., 
    construct=construct@entry=js::NO_CONSTRUCT) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:535
---Type <return> to continue, or q <return> to quit---
#10 0x00007f39e798b78d in InternalCall (cx=0x7f39dea24000, args=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:614
#11 0x00007f39e797e074 in js::CallFromStack (args=..., cx=<optimized out>)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:620
#12 Interpret (cx=0x7f39dea24000, state=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:3462
#13 0x00007f39e798a9d2 in js::RunScript (cx=0x7f39dea24000, state=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:447
#14 0x00007f39e798b20b in js::InternalCallOrConstruct (cx=<optimized out>, cx@entry=0x7f39dea24000, args=..., 
    construct=construct@entry=js::NO_CONSTRUCT) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:587
#15 0x00007f39e798b78d in InternalCall (cx=0x7f39dea24000, args=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:614
#16 0x00007f39e798b900 in js::Call (cx=<optimized out>, fval=..., fval@entry=..., thisv=..., thisv@entry=..., args=..., rval=...)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/vm/Interpreter.cpp:633
#17 0x00007f39e73cbc62 in js::jit::InvokeFunction (cx=<optimized out>, cx@entry=0x7f39dea24000, obj=..., obj@entry=..., 
    constructing=constructing@entry=false, ignoresReturnValue=ignoresReturnValue@entry=false, argc=<optimized out>, argv=0x7ffd422bff70, 
    rval=...) at /home/dinosaur/Downloads/firefox-64.0/js/src/jit/VMFunctions.cpp:112
#18 0x00007f39e73cbf46 in js::jit::InvokeFromInterpreterStub (cx=0x7f39dea24000, frame=0x7ffd422bff48)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/jit/VMFunctions.cpp:142
#19 0x000013810e7c336f in ?? ()
#20 0xfff8800000000001 in ?? ()
#21 0x00007ffd422bff48 in ?? ()
#22 0x00007f39cec66090 in ?? ()
#23 0x00000000000000fc in ?? ()
#24 0x000013810e913bb1 in ?? ()
#25 0x0000000000002842 in ?? ()
#26 0x00007f39d3d18560 in ?? ()
---Type <return> to continue, or q <return> to quit---
#27 0x0000000000000001 in ?? ()
#28 0xfffe7f39da6baf60 in ?? ()
#29 0xfffe7f39da6dfe40 in ?? ()
#30 0x000013810e88ebaa in ?? ()
#31 0x00007ffd422c0038 in ?? ()
#32 0x00007f39d390b318 in ?? ()
#33 0x000013810f4b5873 in ?? ()
#34 0x0000000000009821 in ?? ()
#35 0xfffe7f39da6dfe40 in ?? ()
#36 0xfffe7f39da6baf60 in ?? ()
#37 0xfffe7f39d3d18560 in ?? ()
#38 0xfff9800000000000 in ?? ()
#39 0xfffe7f39ca0303a8 in ?? ()
#40 0xfffe7f39dac70300 in ?? ()
#41 0xfff9800000000000 in ?? ()
#42 0xfffe7f39ca030208 in ?? ()
#43 0xfffe7f39dac70300 in ?? ()
#44 0xfffa80000000000b in ?? ()
#45 0xfffe7f39da6dfe40 in ?? ()
#46 0xfffe7f39ca0302b8 in ?? ()
#47 0x00007f39dea24200 in ?? ()
#48 0x00007ffd422c0040 in ?? ()
#49 0x00007f3900000098 in ?? ()
---Type <return> to continue, or q <return> to quit---
#50 0x00007f39ca030188 in ?? ()
#51 0x00007f39dea24ae8 in ?? ()
#52 0x00000004422c0060 in ?? ()
#53 0x00007ffd422c00a0 in ?? ()
#54 0x000013810e7c0ad7 in ?? ()
#55 0x0000000000002043 in ?? ()
#56 0x00007f39d1ac61f0 in ?? ()
#57 0x0000000000000001 in ?? ()
#58 0xfffe7f39da6baf60 in ?? ()
#59 0xfffe7f39ca030148 in ?? ()
#60 0x00007ffd422c0120 in ?? ()
#61 0x00007ffd422c0110 in ?? ()
#62 0x00007ffd422c0130 in ?? ()
#63 0x000013810f4b3650 in ?? ()
#64 0x00007ffd422c0a80 in ?? ()
#65 0x00007f39dea24000 in ?? ()
#66 0x00007ffd422c05e0 in ?? ()
#67 0x00007f39e7251677 in EnterJit (cx=0xfffe7f39d3d18560, state=..., 
    code=0x7f39cec66090 "`\257k\332\071\177\376\377H\001\003\312\071\177\376\377", '\315' <repeats 15 times>, <incomplete sequence \315>)
    at /home/dinosaur/Downloads/firefox-64.0/js/src/jit/Jit.cpp:103
Backtrace stopped: previous frame inner to this frame (corrupt stack?)

```
