# Workload

简介：主要实现client-go的相关实现，监听k8s apiserver,从队列中获取事件。将事件缓存到cacheMap中。动态展示给前端k8s动态信息。

### 注意点

- prefix: 主要使用前缀区分资源类型。(获取资源前缀名)
- controller: 主要是获取资源。（Get）,如果是添加，更新，修改则调用api实现。查看使用informer
- controller-func: 主要是对外提供的详细方法.基本上都是Get方式