(gdb) bt
#0  0x00007f0ca38a92f0 in views::View::OnKeyEvent(ui::KeyEvent*)@plt () at /home/dinosaur/chrom/src/out/Default/./libviews.so
#1  0x00007f0cba294f74 in ui::EventHandler::OnEvent(ui::Event*) (this=0xe911ac8e088, event=0xe911b47b890) at ../../ui/events/event_handler.cc:31
#2  0x00007f0cba29277a in ui::EventDispatcher::DispatchEvent(ui::EventHandler*, ui::Event*) (this=0x7fff2311ac38, handler=0xe911ac8e088, event=0xe911b47b890) at ../../ui/events/event_dispatcher.cc:193
#3  0x00007f0cba2910ad in ui::EventDispatcher::ProcessEvent(ui::EventTarget*, ui::Event*) (this=0x7fff2311ac38, target=0xe911ac8e048, event=0xe911b47b890) at ../../ui/events/event_dispatcher.cc:142
#4  0x00007f0cba290ad4 in ui::EventDispatcherDelegate::DispatchEventToTarget(ui::EventTarget*, ui::Event*) (this=0xe911a5a83e8, target=0xe911ac8e048, event=0xe911b47b890)
    at ../../ui/events/event_dispatcher.cc:86
#5  0x00007f0cba290990 in ui::EventDispatcherDelegate::DispatchEvent(ui::EventTarget*, ui::Event*) (this=0xe911a5a83e8, target=0xe911ac8e048, event=0xe911b47b890)
    at ../../ui/events/event_dispatcher.cc:58
#6  0x00007f0cba2966d9 in ui::EventProcessor::OnEventFromSource(ui::Event*) (this=0xe911a5a83e8, event=0x7fff23120540) at ../../ui/events/event_processor.cc:57
#7  0x00007f0cba296a1c in non-virtual thunk to ui::EventProcessor::OnEventFromSource(ui::Event*) () at /home/dinosaur/chrom/src/out/Default/./libevents.so
#8  0x00007f0cba298f43 in ui::EventSource::DeliverEventToSink(ui::Event*) (this=0xe911aac6428, event=0x7fff23120540) at ../../ui/events/event_source.cc:106
#9  0x00007f0cba2985a5 in ui::EventSource::SendEventToSinkFromRewriter(ui::Event*, ui::EventRewriter const*) (this=0xe911aac6428, event=0x7fff23120540, rewriter=0x0) at ../../ui/events/event_source.cc:84
#10 0x00007f0cba297a51 in ui::EventSource::SendEventToSink(ui::Event*) (this=0xe911aac6428, event=0x7fff23120540) at ../../ui/events/event_source.cc:43
#11 0x00007f0ca34f28ab in views::Widget::OnKeyEvent(ui::KeyEvent*) (this=0xe911aac6420, event=0x7fff23120540) at ../../ui/views/widget/widget.cc:1182
#12 0x00007f0ca353175b in views::DesktopNativeWidgetAura::OnKeyEvent(ui::KeyEvent*) (this=0xe911a4f3020, event=0x7fff23120540) at ../../ui/views/widget/desktop_aura/desktop_native_widget_aura.cc:1079
#13 0x00007f0cba294f74 in ui::EventHandler::OnEvent(ui::Event*) (this=0xe911a4f3028, event=0x7fff23120540) at ../../ui/events/event_handler.cc:31
#14 0x00007f0cba29277a in ui::EventDispatcher::DispatchEvent(ui::EventHandler*, ui::Event*) (this=0x7fff2311c8e8, handler=0xe911a4f3028, event=0x7fff23120540) at ../../ui/events/event_dispatcher.cc:193
#15 0x00007f0cba2910ad in ui::EventDispatcher::ProcessEvent(ui::EventTarget*, ui::Event*) (this=0x7fff2311c8e8, target=0xe911abee040, event=0x7fff23120540) at ../../ui/events/event_dispatcher.cc:142
#16 0x00007f0cba290ad4 in ui::EventDispatcherDelegate::DispatchEventToTarget(ui::EventTarget*, ui::Event*) (this=0xe911abeec60, target=0xe911abee040, event=0x7fff23120540)
    at ../../ui/events/event_dispatcher.cc:86
#17 0x00007f0cba290990 in ui::EventDispatcherDelegate::DispatchEvent(ui::EventTarget*, ui::Event*) (this=0xe911abeec60, target=0xe911abee040, event=0x7fff23120540)
    at ../../ui/events/event_dispatcher.cc:58
#18 0x00007f0cba2966d9 in ui::EventProcessor::OnEventFromSource(ui::Event*) (this=0xe911abeec60, event=0x7fff23120540) at ../../ui/events/event_processor.cc:57
#19 0x00007f0cba296a1c in non-virtual thunk to ui::EventProcessor::OnEventFromSource(ui::Event*) () at /home/dinosaur/chrom/src/out/Default/./libevents.so
#20 0x00007f0cb746d2d5 in aura::WindowTreeHost::DispatchKeyEventPostIME(ui::KeyEvent*, base::OnceCallback<void (bool)>) (this=0xe911abf1428, event=0x7fff23120540, ack_callback=...)
    at ../../ui/aura/window_tree_host.cc:253
#21 0x00007f0cba244483 in ui::InputMethodBase::DispatchKeyEventPostIME(ui::KeyEvent*, base::OnceCallback<void (bool)>) const (this=0xe911ac41020, event=0x7fff23120540, ack_callback=...)
    at ../../ui/base/ime/input_method_base.cc:168
#22 0x00007f0cba2534fa in ui::InputMethodAuraLinux::ProcessKeyEventDone(ui::KeyEvent*, bool, bool) (this=0xe911ac41020, event=0x7fff23120540, filtered=false, is_handled=false)
    at ../../ui/base/ime/input_method_auralinux.cc:199
#23 0x00007f0cba25256b in ui::InputMethodAuraLinux::DispatchKeyEvent(ui::KeyEvent*) (this=0xe911ac41020, event=0x7fff23120540) at ../../ui/base/ime/input_method_auralinux.cc:109
#24 0x00007f0cb7455ec3 in aura::WindowEventDispatcher::PreDispatchKeyEvent(ui::KeyEvent*) (this=0xe911abeec60, event=0x7fff23120540) at ../../ui/aura/window_event_dispatcher.cc:1096
#25 0x00007f0cb745507a in aura::WindowEventDispatcher::PreDispatchEvent(ui::EventTarget*, ui::Event*) (this=0xe911abeec60, target=0xe911abee040, event=0x7fff23120540)
    at ../../ui/aura/window_event_dispatcher.cc:617
#26 0x00007f0cba290941 in ui::EventDispatcherDelegate::DispatchEvent(ui::EventTarget*, ui::Event*) (this=0xe911abeec60, target=0xe911abee040, event=0x7fff23120540)
    at ../../ui/events/event_dispatcher.cc:54
#27 0x00007f0cba2966d9 in ui::EventProcessor::OnEventFromSource(ui::Event*) (this=0xe911abeec60, event=0x7fff23120540) at ../../ui/events/event_processor.cc:57
#28 0x00007f0cba296a1c in non-virtual thunk to ui::EventProcessor::OnEventFromSource(ui::Event*) () at /home/dinosaur/chrom/src/out/Default/./libevents.so
#29 0x00007f0cba298f43 in ui::EventSource::DeliverEventToSink(ui::Event*) (this=0xe911abf1430, event=0x7fff23120540) at ../../ui/events/event_source.cc:106
#30 0x00007f0cba2985a5 in ui::EventSource::SendEventToSinkFromRewriter(ui::Event*, ui::EventRewriter const*) (this=0xe911abf1430, event=0x7fff23120540, rewriter=0x0) at ../../ui/events/event_source.cc:84
#31 0x00007f0cba297a51 in ui::EventSource::SendEventToSink(ui::Event*) (this=0xe911abf1430, event=0x7fff23120540) at ../../ui/events/event_source.cc:43
#32 0x00007f0ca355c4c8 in views::DesktopWindowTreeHostX11::DispatchKeyEvent(ui::KeyEvent*) (this=0xe911abf1420, event=0x7fff23120540)
    at ../../ui/views/widget/desktop_aura/desktop_window_tree_host_x11.cc:1865
#33 0x00007f0ca355e08e in views::DesktopWindowTreeHostX11::DispatchEvent(_XEvent* const&) (this=0xe911abf1420, event=@0x7fff23121050: 0x7fff231210e8)
    at ../../ui/views/widget/desktop_aura/desktop_window_tree_host_x11.cc:2051
#34 0x00007f0cbd2c1a27 in ui::PlatformEventSource::DispatchEvent(_XEvent*) (this=0xe911a4848e8, platform_event=0x7fff231210e8) at ../../ui/events/platform/platform_event_source.cc:91
#35 0x00007f0c9d496811 in ui::X11EventSourceGlib::ProcessXEvent(_XEvent*) (this=0xe911a4848e0, xevent=0x7fff231210e8) at ../../ui/events/platform/x11/x11_event_source_glib.cc:64
#36 0x00007f0c9d484dc6 in ui::X11EventSource::ExtractCookieDataDispatchEvent(_XEvent*) (this=0xe911a4849d0, xevent=0x7fff231210e8) at ../../ui/events/platform/x11/x11_event_source.cc:248
#37 0x00007f0c9d484d34 in ui::X11EventSource::DispatchXEvents() (this=0xe911a4849d0) at ../../ui/events/platform/x11/x11_event_source.cc:141
#38 0x00007f0c9d496cd5 in ui::(anonymous namespace)::XSourceDispatch(_GSource*, int (*)(void*), void*) (source=0xe911a487180, unused_func=0x0, data=0xe911a4849d0)
    at ../../ui/events/platform/x11/x11_event_source_glib.cc:40
#39 0x00007f0ca17aa0f5 in g_main_context_dispatch () at /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0
#40 0x00007f0ca17aa4c0 in  () at /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0
#41 0x00007f0ca17aa54c in g_main_context_iteration () at /usr/lib/x86_64-linux-gnu/libglib-2.0.so.0
---Type <return> to continue, or q <return> to quit---
#42 0x00007f0cc58e929f in base::MessagePumpGlib::Run(base::MessagePump::Delegate*) (this=0xe911a54af00, delegate=0xe911a366720) at ../../base/message_loop/message_pump_glib.cc:305
#43 0x00007f0cc58e08fb in base::MessageLoop::Run(bool) (this=0xe911a366720, application_tasks_allowed=true) at ../../base/message_loop/message_loop.cc:502
#44 0x00007f0cc598d89d in base::RunLoop::Run() (this=0x7fff23121bc0) at ../../base/run_loop.cc:102
#45 0x00007f0cc1a1f34c in content::BrowserMainLoop::MainMessageLoopRun() (this=0xe911a971da0) at ../../content/browser/browser_main_loop.cc:1524
#46 0x00007f0cc1a1efa2 in content::BrowserMainLoop::RunMainMessageLoopParts() (this=0xe911a971da0) at ../../content/browser/browser_main_loop.cc:996
#47 0x00007f0cc1a275a0 in content::BrowserMainRunnerImpl::Run() (this=0xe911a964c80) at ../../content/browser/browser_main_runner_impl.cc:165
#48 0x00000000019d4274 in ShellBrowserMain(content::MainFunctionParams const&, std::__1::unique_ptr<content::BrowserMainRunner, std::__1::default_delete<content::BrowserMainRunner> > const&) (parameters=..., main_runner=...) at ../../content/shell/browser/shell_browser_main.cc:30
#49 0x000000000190d5ee in content::ShellMainDelegate::RunProcess(std::__1::basic_string<char, std::__1::char_traits<char>, std::__1::allocator<char> > const&, content::MainFunctionParams const&) (this=
    0x7fff23123af8, process_type=..., main_function_params=...) at ../../content/shell/app/shell_main_delegate.cc:371
#50 0x00007f0cc3a364ea in content::RunBrowserProcessMain(content::MainFunctionParams const&, content::ContentMainDelegate*) (main_function_params=..., delegate=0x7fff23123af8)
    at ../../content/app/content_main_runner_impl.cc:527
#51 0x00007f0cc3a39256 in content::ContentMainRunnerImpl::Run(bool) (this=0xe911a34dbd0, start_service_manager_only=false) at ../../content/app/content_main_runner_impl.cc:902
#52 0x00007f0cc3a2f24c in content::ContentServiceManagerMainDelegate::RunEmbedderProcess() (this=0x7fff23123a68) at ../../content/app/content_service_manager_main_delegate.cc:53
#53 0x00007f0c9d4b9c2a in service_manager::Main(service_manager::MainParams const&) (params=...) at ../../services/service_manager/embedder/main.cc:472
#54 0x00007f0cc3a33c43 in content::ContentMain(content::ContentMainParams const&) (params=...) at ../../content/app/content_main.cc:19
#55 0x000000000075b165 in main(int, char const**) (argc=3, argv=0x7fff23123c28) at ../../content/shell/app/shell_main.cc:39
```
```
