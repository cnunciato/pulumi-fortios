# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_fortios.filter.dns as __dns
    dns = __dns
    import pulumiverse_fortios.filter.email as __email
    email = __email
    import pulumiverse_fortios.filter.file as __file
    file = __file
    import pulumiverse_fortios.filter.sctp as __sctp
    sctp = __sctp
    import pulumiverse_fortios.filter.spam as __spam
    spam = __spam
    import pulumiverse_fortios.filter.ssh as __ssh
    ssh = __ssh
    import pulumiverse_fortios.filter.video as __video
    video = __video
    import pulumiverse_fortios.filter.web as __web
    web = __web
else:
    dns = _utilities.lazy_import('pulumiverse_fortios.filter.dns')
    email = _utilities.lazy_import('pulumiverse_fortios.filter.email')
    file = _utilities.lazy_import('pulumiverse_fortios.filter.file')
    sctp = _utilities.lazy_import('pulumiverse_fortios.filter.sctp')
    spam = _utilities.lazy_import('pulumiverse_fortios.filter.spam')
    ssh = _utilities.lazy_import('pulumiverse_fortios.filter.ssh')
    video = _utilities.lazy_import('pulumiverse_fortios.filter.video')
    web = _utilities.lazy_import('pulumiverse_fortios.filter.web')
