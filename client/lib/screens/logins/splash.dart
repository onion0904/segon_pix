import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import '../../logic/db/user_manager.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import '../../logic/http/get.dart';
import '../../model/provider.dart';

class Splash extends HookConsumerWidget {
  const Splash({super.key, required this.changeIndex});

  final void Function(int) changeIndex;

  @override
  Widget build(context, ref) {
    final isCompleteLoad = useState<int?>(null);

    Future<void> init() async {
      final result = await UserManager.initializeUserManager();
      isCompleteLoad.value = result;
    }

    useEffect(() {
      init();
      return null;
    }, []);

    if (isCompleteLoad.value == null) {
      return Center(
        child: Image.network(
            "https://onion0904.dev/ocGvg5tH5gfqsDS1715839141_1715839204.png"), // スプラッシュ画面の画像
      );
    } else if (isCompleteLoad.value == 1) {

      WidgetsBinding.instance.addPostFrameCallback((_) async {
        ref.read(userProvider.notifier).state =
            await getUserWithAuth(token: UserManager.token, email: UserManager.email, password: UserManager.password);
        if (context.mounted) {
          context.go("/hub");
        }
      });
      return const Center();
    } else {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        changeIndex(1);
      });
      return const Center();
    }
  }
}
